// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"net/netip"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/command"
	"github.com/cilium/cilium/pkg/common"
	"github.com/cilium/cilium/pkg/egressgateway"
	"github.com/cilium/cilium/pkg/maps/egressmap"
)

const (
	egressListUsage = "List egress policy entries."
)

type egressPolicy struct {
	SourceIP  string
	DestCIDR  string
	EgressIP  string
	GatewayIP string
}

var bpfEgressListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List egress policy entries",
	Long:    egressListUsage,
	Run: func(cmd *cobra.Command, args []string) {
		common.RequireRootPrivilege("cilium bpf egress list")

		bpfEgressList := []egressPolicy{}
		var ipv4MapExists, ipv6MapExists bool

		policyMap4, err := egressmap.OpenPinnedPolicyMap4(log)
		if err == nil {
			ipv4MapExists = true
			parse4 := func(key *egressmap.EgressPolicyKey4, val *egressmap.EgressPolicyVal4) {
				bpfEgressList = append(bpfEgressList, egressPolicy{
					SourceIP:  key.GetSourceIP().String(),
					DestCIDR:  key.GetDestCIDR().String(),
					EgressIP:  val.GetEgressAddr().String(),
					GatewayIP: mapGatewayIP(val.GetGatewayAddr()),
				})
			}

			if err := policyMap4.IterateWithCallback(parse4); err != nil {
				Fatalf("Error dumping contents of IPv4 egress policy map: %s\n", err)
			}
		} else if !errors.Is(err, fs.ErrNotExist) {
			Fatalf("Cannot open IPv4 egress gateway bpf map: %s", err)
		}

		policyMap6, err := egressmap.OpenPinnedPolicyMap6(log)
		if err == nil {
			ipv6MapExists = true
			parse6 := func(key *egressmap.EgressPolicyKey6, val *egressmap.EgressPolicyVal6) {
				bpfEgressList = append(bpfEgressList, egressPolicy{
					SourceIP:  key.GetSourceIP().String(),
					DestCIDR:  key.GetDestCIDR().String(),
					EgressIP:  val.GetEgressAddr().String(),
					GatewayIP: mapGatewayIP(val.GetGatewayAddr()),
				})
			}

			if err := policyMap6.IterateWithCallback(parse6); err != nil {
				Fatalf("Error dumping contents of IPv6 egress policy map: %s\n", err)
			}
		} else if !errors.Is(err, fs.ErrNotExist) {
			Fatalf("Cannot open IPv6 egress gateway bpf map: %s", err)
		}

		if !ipv4MapExists && !ipv6MapExists {
			fmt.Fprintln(os.Stderr, "Cannot find egress gateway bpf maps")
			return
		}

		if command.OutputOption() {
			if err := command.PrintOutput(bpfEgressList); err != nil {
				Fatalf("error getting output of map in %s: %s\n", command.OutputOptionString(), err)
			}
			return
		}

		if len(bpfEgressList) == 0 {
			fmt.Fprintf(os.Stderr, "No entries found.\n")
		} else {
			printEgressList(bpfEgressList)
		}
	},
}

// This function attempt to translate gatewayIP to special values if they exist
// or return the IP as a string otherwise.
func mapGatewayIP(ip netip.Addr) string {
	if ip == egressgateway.GatewayNotFoundIPv4 {
		return "Not Found"
	}
	if ip == egressgateway.ExcludedCIDRIPv4 {
		return "Excluded CIDR"
	}
	return ip.String()
}

func printEgressList(egressList []egressPolicy) {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)

	fmt.Fprintln(w, "Source IP\tDestination CIDR\tEgress IP\tGateway IP")
	for _, ep := range egressList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", ep.SourceIP, ep.DestCIDR, ep.EgressIP, ep.GatewayIP)
	}

	w.Flush()
}

func init() {
	BPFEgressCmd.AddCommand(bpfEgressListCmd)
	command.AddOutputOption(bpfEgressListCmd)
}
