// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nomad

import (
	"github.com/armon/go-metrics"

	"github.com/hashicorp/nomad/nomad/structs"
)

// MeasureRPCRate increments the appropriate rate metric for this endpoint,
// with a label from the identity
func (srv *Server) MeasureRPCRate(endpoint, op string, args structs.RequestWithIdentity) {
	identity := args.GetIdentity()

	if !srv.config.ACLEnabled || identity == nil || srv.config.DisableRPCRateMetricsLabels {
		// If ACLs aren't enabled, we never have a sensible identity.
		// Or the administrator may have disabled the identity labels.
		metrics.IncrCounter([]string{"nomad", "rpc", endpoint, op}, 1)
	} else {
		metrics.IncrCounterWithLabels(
			[]string{"nomad", "rpc", endpoint, op}, 1,
			[]metrics.Label{{Name: "identity", Value: identity.String()}})
	}
}
