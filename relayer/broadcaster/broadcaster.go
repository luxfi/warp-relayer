// SPDX-License-Identifier: Lux-Ecosystem-1.2
// Copyright (c) 2020-2025 Lux Industries Inc.

// Package broadcaster builds + submits the destination L1 tx that
// includes the aggregated warp message in its access list, so the
// destination's getVerifiedWarpMessage precompile returns valid=true.
//
// Status: scaffold. Real builder + signer + submitter logic lands in
// follow-up commits.
package broadcaster

// Broadcaster submits to one destination L1.
type Broadcaster struct {
	// Fields land in follow-up.
}
