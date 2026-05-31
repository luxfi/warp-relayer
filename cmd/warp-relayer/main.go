// SPDX-License-Identifier: Lux-Ecosystem-1.2
// Copyright (c) 2020-2025 Lux Industries Inc.

// Command warp-relayer is the Lux-native off-chain courier for
// cross-chain warp messages between Lux-derived L1s. See README.md for
// scope (asset-bridge teleport; NOT validator-set management).
//
// Status: scaffold. Real subscriber/aggregator/broadcaster logic lands
// in follow-up commits as the first bridge consumer (likely ICTT
// TokenHome/TokenRemote) comes online.
package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `warp-relayer — cross-chain Lux warp message courier.

Usage:
  warp-relayer --config /path/to/config.json

Status: scaffold; full implementation pending first bridge consumer.
`

func main() {
	configPath := flag.String("config", "", "path to relayer config JSON")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	if *showVersion {
		fmt.Println("warp-relayer (scaffold) — see luxfi/warp-relayer README.md")
		return
	}

	if *configPath == "" {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Fprintln(os.Stderr,
		"warp-relayer: scaffold only — subscriber/aggregator/broadcaster not yet implemented.")
	fmt.Fprintln(os.Stderr,
		"             config path: "+*configPath)
	os.Exit(1)
}
