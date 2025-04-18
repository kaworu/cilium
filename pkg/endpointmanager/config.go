// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package endpointmanager

import (
	"github.com/spf13/pflag"

	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/time"
)

type EndpointManagerConfig struct {
	// EndpointGCInterval is interval to attempt garbage collection of
	// endpoints that are no longer alive and healthy.
	EndpointGCInterval time.Duration

	// EndpointRegenInterval is interval between periodic endpoint regenerations.
	EndpointRegenInterval time.Duration
}

func (def EndpointManagerConfig) Flags(flags *pflag.FlagSet) {
	flags.Duration(option.EndpointGCInterval, def.EndpointGCInterval,
		"Periodically monitor local endpoint health via link status on this interval and garbage collect them if they become unhealthy, set to 0 to disable")
	flags.MarkHidden(option.EndpointGCInterval)

	flags.Duration(option.EndpointRegenInterval, def.EndpointRegenInterval,
		"Periodically recalculate and re-apply endpoint configuration. Set to 0 to disable")
}

var defaultEndpointManagerConfig = EndpointManagerConfig{
	EndpointGCInterval:    5 * time.Minute,
	EndpointRegenInterval: 2 * time.Minute,
}
