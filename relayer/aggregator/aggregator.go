// SPDX-License-Identifier: Lux-Ecosystem-1.2
// Copyright (c) 2020-2025 Lux Industries Inc.

// Package aggregator collects BLS signature shares from the source L1's
// validator set and aggregates them into a single 48-byte BLS aggregate
// the destination chain's warp precompile can verify.
//
// Status: scaffold. Real aggregation logic lands in follow-up commits;
// it will compose against luxfi/warp's signer + signature primitives.
package aggregator

// Aggregator collects BLS shares + aggregates.
type Aggregator struct {
	// Fields land in follow-up.
}
