// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressgateway

import (
	"fmt"
	"hash/fnv"
	"log/slog"
	"net/netip"
	"slices"

	"go4.org/netipx"
	"k8s.io/apimachinery/pkg/types"

	"github.com/cilium/cilium/pkg/datapath/linux/netdevice"
	"github.com/cilium/cilium/pkg/datapath/linux/route"
	k8sConst "github.com/cilium/cilium/pkg/k8s/apis/cilium.io"
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	k8sLabels "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/labels"
	slim_metav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	"github.com/cilium/cilium/pkg/logging/logfields"
	nodeTypes "github.com/cilium/cilium/pkg/node/types"
	"github.com/cilium/cilium/pkg/policy"
	"github.com/cilium/cilium/pkg/policy/api"
)

// policyGatewayConfig is the internal representation of an egress gateway,
// describing which node should act as egress gateway for a given policy.
type policyGatewayConfig struct {
	nodeSelector api.EndpointSelector
	iface        string
	egressIP     netip.Addr
}

// gatewayConfig is the gateway configuration derived at runtime from a policy.
//
// Some of these fields are derived from the running system as the policy may
// specify only the egress IP (and so we need to figure out which interface has
// that IP assigned to) or the interface (and in this case we need to find the
// first IPv4 assigned to that).
type gatewayConfig struct {
	// ifaceName is the name of the interface used to SNAT traffic
	ifaceName string
	// egressIP4 is the IP used to SNAT traffic with IPv4 policies
	egressIP4 netip.Addr
	// egressIP6 is the IP used to SNAT traffic with IPv6 policies
	egressIP6 netip.Addr
	// gatewayIP is the node internal IP of the gateway
	gatewayIP netip.Addr
	// localNodeConfiguredAsGateway tells if the local node is configured to
	// act as an egress gateway node for this config.
	// This information is used to decide if it is necessary to relax the rp_filter
	// on the interface used to SNAT traffic
	localNodeConfiguredAsGateway bool
}

// PolicyConfig is the internal representation of CiliumEgressGatewayPolicy.
type PolicyConfig struct {
	// id is the parsed config name and namespace
	id types.NamespacedName

	endpointSelectors []api.EndpointSelector
	nodeSelectors     []api.EndpointSelector
	dstCIDRs          []netip.Prefix
	excludedCIDRs     []netip.Prefix
	policyGwConfigs   []policyGatewayConfig
	gatewayConfigs    []gatewayConfig
	matchedEndpoints  map[endpointID]*endpointMetadata
	v6Needed          bool
}

// PolicyID includes policy name and namespace
type policyID = types.NamespacedName

// matchesEndpointLabels determines if the given endpoint is a match for the
// policy config based on matching labels.
func (config *PolicyConfig) matchesEndpointLabels(endpointInfo *endpointMetadata) bool {
	labelsToMatch := k8sLabels.Set(endpointInfo.labels)
	for _, selector := range config.endpointSelectors {
		if selector.Matches(labelsToMatch) {
			return true
		}
	}
	return false
}

// matchesNodeLabels determines if the given node lables is a match for the
// policy config based on matching labels.
func (config *PolicyConfig) matchesNodeLabels(nodeLabels map[string]string) bool {
	if len(config.nodeSelectors) == 0 {
		return true
	}
	labelsToMatch := k8sLabels.Set(nodeLabels)
	for _, selector := range config.nodeSelectors {
		if selector.Matches(labelsToMatch) {
			return true
		}
	}
	return false
}

// updateMatchedEndpointIDs update the policy's cache of matched endpoint IDs
func (config *PolicyConfig) updateMatchedEndpointIDs(epDataStore map[endpointID]*endpointMetadata, nodesAddresses2Labels map[string]map[string]string) {
	config.matchedEndpoints = make(map[endpointID]*endpointMetadata)
	for _, endpoint := range epDataStore {
		if config.matchesEndpointLabels(endpoint) && config.matchesNodeLabels(nodesAddresses2Labels[endpoint.nodeIP]) {
			config.matchedEndpoints[endpoint.id] = endpoint
		}
	}
}

func (config *policyGatewayConfig) selectsNodeAsGateway(node nodeTypes.Node) bool {
	return config.nodeSelector.Matches(k8sLabels.Set(node.Labels))
}

func (config *PolicyConfig) regenerateGatewayConfig(manager *Manager) {
	config.gatewayConfigs = make([]gatewayConfig, 0, len(config.policyGwConfigs))

	for _, policyGwc := range config.policyGwConfigs {
		gwc := gatewayConfig{
			egressIP4: netip.IPv4Unspecified(),
			egressIP6: netip.IPv6Unspecified(),
			gatewayIP: GatewayNotFoundIPv4,
		}

		for _, node := range manager.nodes {
			if !policyGwc.selectsNodeAsGateway(node) {
				continue
			}

			addr, ok := netipx.FromStdIP(node.GetNodeIP(false))
			if !ok {
				continue
			}
			gwc.gatewayIP = addr

			if node.IsLocal() {
				err := gwc.deriveFromPolicyGatewayConfig(manager.logger, &policyGwc, config.v6Needed)
				if err != nil {
					manager.logger.Error(
						"Failed to derive policy gateway configuration",
						logfields.Error, err,
						logfields.CiliumEgressGatewayPolicyName, config.id,
						logfields.Interface, policyGwc.iface,
						logfields.EgressIP, policyGwc.egressIP,
					)
				}
			}

			break
		}

		config.gatewayConfigs = append(config.gatewayConfigs, gwc)
	}
}

// deriveFromPolicyGatewayConfig retrieves all the missing gateway configuration
// data (such as egress IP or interface) given a policy egress gateway config
func (gwc *gatewayConfig) deriveFromPolicyGatewayConfig(logger *slog.Logger, gc *policyGatewayConfig, v6Needed bool) error {
	var err error
	var egressIP4, egressIP6 netip.Addr

	gwc.localNodeConfiguredAsGateway = false
	gwc.egressIP4 = EgressIPNotFoundIPv4
	gwc.egressIP6 = EgressIPNotFoundIPv6

	switch {
	case gc.iface != "":
		// If the gateway config specifies an interface, use the first IPv4/v6 assigned to that
		// interface as egress IPs
		gwc.ifaceName = gc.iface
		egressIP4, err = netdevice.GetIfaceFirstIPv4Address(gc.iface)
		if err != nil {
			return fmt.Errorf("failed to retrieve IPv4 address for egress interface: %w", err)
		}

		if v6Needed {
			egressIP6, err = netdevice.GetIfaceFirstIPv6Address(gc.iface)
			if err != nil {
				return fmt.Errorf("failed to retrieve IPv6 address for egress interface: %w", err)
			}
		}
	case gc.egressIP.IsValid():
		// If the gateway config specifies an egress IP, use the interface with that IP as egress
		// interface.
		egressIP4 = gc.egressIP

		// TODO: add ipv6 support for specifying an egress IP, currently only ipv4 is supported.
		if v6Needed {
			egressIP6 = EgressIPNotFoundIPv6
		}

		gwc.ifaceName, err = netdevice.GetIfaceWithIPv4Address(gc.egressIP)
		if err != nil {
			return fmt.Errorf("failed to retrieve interface with egress IP: %w", err)
		}

	default:
		// If the gateway config doesn't specify any egress IP or interface, use the
		// interface with the IPv4 default route
		iface, err := route.NodeDeviceWithDefaultRoute(logger, true, false)
		if err != nil {
			return fmt.Errorf("failed to find interface with IPv4 default route: %w", err)
		}

		gwc.ifaceName = iface.Attrs().Name
		egressIP4, err = netdevice.GetIfaceFirstIPv4Address(gwc.ifaceName)
		if err != nil {
			return fmt.Errorf("failed to retrieve IPv4 address for egress interface: %w", err)
		}

		if v6Needed {
			iface, err := route.NodeDeviceWithDefaultRoute(logger, false, true)
			if err != nil {
				return fmt.Errorf("failed to find interface with IPv6 default route: %w", err)
			}

			if iface.Attrs().Name != gwc.ifaceName {
				return fmt.Errorf("IPv6 default route interface doesn't match IPv4 default route interface")
			}

			gwc.ifaceName = iface.Attrs().Name
			egressIP6, err = netdevice.GetIfaceFirstIPv6Address(gwc.ifaceName)
			if err != nil {
				return fmt.Errorf("failed to retrieve IPv6 address for egress interface: %w", err)
			}
		}
	}

	gwc.egressIP4 = egressIP4
	if v6Needed {
		gwc.egressIP6 = egressIP6
	}
	gwc.localNodeConfiguredAsGateway = true

	return nil
}

func computeEndpointHash(endpointUID types.UID) uint32 {
	h := fnv.New32a()
	h.Write([]byte(endpointUID))
	return h.Sum32()
}

// forEachEndpointAndCIDR iterates through each combination of endpoints and
// destination/excluded CIDRs of the receiver policy, and for each of them it
// calls the f callback function passing the given endpoint and CIDR, together
// with a boolean value indicating if the CIDR belongs to the excluded ones and
// the gatewayConfig of the receiver policy.
// For multigateway policies the gateways are ordered by IP and paired with each
// endpoint using the hash of the endpoint UID.
func (config *PolicyConfig) forEachEndpointAndCIDR(f func(netip.Addr, netip.Prefix, bool, *gatewayConfig)) {
	// Sort gateways to get consistent assignments across nodes.
	slices.SortFunc(config.gatewayConfigs, func(a, b gatewayConfig) int {
		return a.gatewayIP.Compare(b.gatewayIP)
	})

	for _, endpoint := range config.matchedEndpoints {
		var gateway *gatewayConfig
		if len(config.gatewayConfigs) > 1 {
			index := computeEndpointHash(endpoint.id) % uint32(len(config.gatewayConfigs))
			gateway = &config.gatewayConfigs[index]
		} else {
			gateway = &config.gatewayConfigs[0]
		}

		for _, endpointIP := range endpoint.ips {
			isExcludedCIDR := false
			for _, dstCIDR := range config.dstCIDRs {
				f(endpointIP, dstCIDR, isExcludedCIDR, gateway)
			}

			isExcludedCIDR = true
			for _, excludedCIDR := range config.excludedCIDRs {
				f(endpointIP, excludedCIDR, isExcludedCIDR, gateway)
			}
		}
	}
}

func parseEgressGateway(egressGateway *v2.EgressGateway) (*policyGatewayConfig, error) {
	if egressGateway == nil {
		return nil, fmt.Errorf("egressGateway can't be empty")
	}

	if egressGateway.Interface != "" && egressGateway.EgressIP != "" {
		return nil, fmt.Errorf("gateway configuration can't specify both an interface and an egress IP")
	}

	policyGwc := &policyGatewayConfig{
		nodeSelector: api.NewESFromK8sLabelSelector("", egressGateway.NodeSelector),
		iface:        egressGateway.Interface,
	}

	// EgressIP is not a required field, validate and parse it only if non-empty
	if egressGateway.EgressIP != "" {
		addr, err := netip.ParseAddr(egressGateway.EgressIP)
		if err != nil {
			return nil, fmt.Errorf("failed to parse egress IP %s: %w", egressGateway.EgressIP, err)
		}

		policyGwc.egressIP = addr
	}

	return policyGwc, nil
}

// ParseCEGP takes a CiliumEgressGatewayPolicy CR and converts to PolicyConfig,
// the internal representation of the egress gateway policy
func ParseCEGP(cegp *v2.CiliumEgressGatewayPolicy) (*PolicyConfig, error) {
	var endpointSelectorList []api.EndpointSelector
	var nodeSelectorList []api.EndpointSelector
	var dstCidrList []netip.Prefix
	var excludedCIDRs []netip.Prefix
	var policyGwConfigs []policyGatewayConfig
	var v6Needed bool

	allowAllNamespacesRequirement := slim_metav1.LabelSelectorRequirement{
		Key:      k8sConst.PodNamespaceLabel,
		Operator: slim_metav1.LabelSelectorOpExists,
	}

	name := cegp.ObjectMeta.Name
	if name == "" {
		return nil, fmt.Errorf("must have a name")
	}

	destinationCIDRs := cegp.Spec.DestinationCIDRs
	if destinationCIDRs == nil {
		return nil, fmt.Errorf("destinationCIDRs can't be empty")
	}

	for _, egressGateway := range cegp.Spec.EgressGateways {
		policyGwc, err := parseEgressGateway(&egressGateway)
		if err != nil {
			return nil, err
		}
		policyGwConfigs = append(policyGwConfigs, *policyGwc)
	}

	// If there are any elements in EgressGateways skip the EgressGateway field.
	if len(policyGwConfigs) == 0 {
		egressGateway := cegp.Spec.EgressGateway
		policyGwc, err := parseEgressGateway(egressGateway)
		if err != nil {
			return nil, err
		}
		policyGwConfigs = append(policyGwConfigs, *policyGwc)
	}

	for _, cidrString := range destinationCIDRs {
		cidr, err := netip.ParsePrefix(string(cidrString))
		if err != nil {
			return nil, fmt.Errorf("failed to parse destination CIDR %s: %w", cidrString, err)
		}
		dstCidrList = append(dstCidrList, cidr)
		if cidr.Addr().Is6() {
			v6Needed = true
		}
	}

	for _, cidrString := range cegp.Spec.ExcludedCIDRs {
		cidr, err := netip.ParsePrefix(string(cidrString))
		if err != nil {
			return nil, fmt.Errorf("failed to parse excluded CIDR %s: %w", cidr, err)
		}
		excludedCIDRs = append(excludedCIDRs, cidr)
	}

	for _, egressRule := range cegp.Spec.Selectors {
		if egressRule.NodeSelector != nil {
			nodeSelectorList = append(
				nodeSelectorList,
				api.NewESFromK8sLabelSelector("", egressRule.NodeSelector))
		}
		if egressRule.NamespaceSelector != nil {
			prefixedNsSelector := egressRule.NamespaceSelector
			matchLabels := map[string]string{}
			// We use our own special label prefix for namespace metadata,
			// thus we need to prefix that prefix to all NamespaceSelector.MatchLabels
			for k, v := range egressRule.NamespaceSelector.MatchLabels {
				matchLabels[policy.JoinPath(k8sConst.PodNamespaceMetaLabels, k)] = v
			}

			prefixedNsSelector.MatchLabels = matchLabels

			// We use our own special label prefix for namespace metadata,
			// thus we need to prefix that prefix to all NamespaceSelector.MatchLabels
			for i, lsr := range egressRule.NamespaceSelector.MatchExpressions {
				lsr.Key = policy.JoinPath(k8sConst.PodNamespaceMetaLabels, lsr.Key)
				prefixedNsSelector.MatchExpressions[i] = lsr
			}

			// Empty namespace selector selects all namespaces (i.e., a namespace
			// label exists).
			if len(egressRule.NamespaceSelector.MatchLabels) == 0 && len(egressRule.NamespaceSelector.MatchExpressions) == 0 {
				prefixedNsSelector.MatchExpressions = []slim_metav1.LabelSelectorRequirement{allowAllNamespacesRequirement}
			}

			endpointSelectorList = append(
				endpointSelectorList,
				api.NewESFromK8sLabelSelector("", prefixedNsSelector, egressRule.PodSelector))
		} else if egressRule.PodSelector != nil {
			endpointSelectorList = append(
				endpointSelectorList,
				api.NewESFromK8sLabelSelector("", egressRule.PodSelector))
		} else {
			return nil, fmt.Errorf("cannot have both nil namespace selector and nil pod selector")
		}
	}

	return &PolicyConfig{
		endpointSelectors: endpointSelectorList,
		nodeSelectors:     nodeSelectorList,
		dstCIDRs:          dstCidrList,
		excludedCIDRs:     excludedCIDRs,
		matchedEndpoints:  make(map[endpointID]*endpointMetadata),
		policyGwConfigs:   policyGwConfigs,
		v6Needed:          v6Needed,
		id: types.NamespacedName{
			Name: name,
		},
	}, nil
}

// ParseCEGPConfigID takes a CiliumEgressGatewayPolicy CR and returns only the config id
func ParseCEGPConfigID(cegp *v2.CiliumEgressGatewayPolicy) types.NamespacedName {
	return policyID{
		Name: cegp.Name,
	}
}
