// SPDX-License-Identifier: Lux-Ecosystem-1.2
// Copyright (c) 2020-2025 Lux Industries Inc.

// Package config holds the per-source + per-destination configuration
// schema for warp-relayer. Wire format follows luxfi/warp (this package
// describes how to drive it, not how it works on the wire).
//
// Status: scaffold. Field set will fill in as the first consumer
// (likely an ICTT TokenHome → TokenRemote bridge) lands.
package config

// Config is the top-level relayer configuration.
type Config struct {
	// LogLevel is one of: debug | info | warn | error.
	LogLevel string `json:"logLevel,omitempty"`

	// NetworkID is the Lux primary network id this relayer runs against
	// (1=mainnet, 2=testnet, 3=devnet, 1337=local). Used to select the
	// default p-chain endpoint + validator-set queries.
	NetworkID uint32 `json:"networkId"`

	// PChainAPIURL is the primary network P-Chain RPC the relayer uses
	// to query each source L1's current validator set + BLS pubkeys.
	PChainAPIURL string `json:"pChainAPIURL"`

	// Sources is the list of source L1s to listen on for SendWarpMessage
	// events. Each Source declares which message-emitting contracts to
	// subscribe to.
	Sources []Source `json:"sources"`

	// Destinations is the list of destination L1s where aggregated warp
	// messages are submitted. Each Destination carries the gas-payer
	// key reference (KMS-projected; never a plaintext env-var EOA).
	Destinations []Destination `json:"destinations"`
}

// Source describes a source L1 to listen on.
type Source struct {
	// ChainID is the blockchain id (CB58) of the source L1.
	ChainID string `json:"chainId"`

	// RPCEndpoint is the EVM JSON-RPC endpoint for state queries.
	RPCEndpoint string `json:"rpcEndpoint"`

	// WSEndpoint is the EVM websocket endpoint for log subscriptions.
	WSEndpoint string `json:"wsEndpoint"`

	// MessageContracts maps message-emitter contract addresses (0x-hex)
	// to a per-protocol descriptor. Common protocols: ictt, teleporter.
	MessageContracts map[string]MessageContractConfig `json:"messageContracts"`
}

// MessageContractConfig describes one warp-emitting contract.
type MessageContractConfig struct {
	// Protocol is the warp-message protocol family (e.g., "ictt",
	// "teleporter"). Drives the relayer's payload unpacking.
	Protocol string `json:"protocol"`

	// Settings is a protocol-specific bag of options. Schema lives in
	// the protocol handler package, not here.
	Settings map[string]any `json:"settings,omitempty"`
}

// Destination describes a destination L1 to submit aggregated messages
// to.
type Destination struct {
	// ChainID is the blockchain id (CB58) of the destination L1, or a
	// special token like "p-chain" for the Lux primary network.
	ChainID string `json:"chainId"`

	// RPCEndpoint is the EVM JSON-RPC endpoint the relayer broadcasts
	// the destination tx to.
	RPCEndpoint string `json:"rpcEndpoint"`

	// AccountKey is a reference to the gas-payer private key. MUST be
	// sourced from Lux KMS — see AccountKeyRef. Plaintext key fields
	// are intentionally absent.
	AccountKey AccountKeyRef `json:"accountKey"`
}

// AccountKeyRef points at the relayer's gas-payer key in Lux KMS.
// Plaintext private-key fields are intentionally not part of this
// schema — relayer pods MUST fetch via the KMS client at startup.
type AccountKeyRef struct {
	// Host is the KMS endpoint, e.g.
	// "https://kms.hanzo.ai" or "http://kms.tenant.svc:9999".
	Host string `json:"host"`

	// Path is the KMS path under which the relayer key lives, e.g.
	// "relayer/devnet/account-private-key".
	Path string `json:"path"`

	// Env is the KMS env scope (e.g. "devnet"); orthogonal to Path
	// per Lux KMS convention.
	Env string `json:"env"`
}
