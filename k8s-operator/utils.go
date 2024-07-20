// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !plan9

// Package kube contains types and utilities for the Tailscale Kubernetes Operator.
package kube

import (
	"fmt"

	"tailscale.com/tailcfg"
)

const (
	Alpha1Version = "v1alpha1"

	DNSRecordsCMName = "dnsrecords"
	DNSRecordsCMKey  = "records.json"
)

type Records struct {
	// Version is the version of this Records configuration. Version is
	// written by the operator, i.e when it first populates the Records.
	// k8s-nameserver must verify that it knows how to parse a given
	// version.
	Version string `json:"version"`
	// IP4 contains a mapping of DNS names to IPv4 address(es).
	IP4 map[string][]string `json:"ip4"`
}

// TailscaledConfigFileNameForCap returns a tailscaled config file name in
// format expected by containerboot for the given CapVer.
func TailscaledConfigFileNameForCap(cap tailcfg.CapabilityVersion) string {
	if cap < 95 {
		return "tailscaled"
	}
	return fmt.Sprintf("cap-%v.hujson", cap)
}

// CapVerFromFileName parses the capability version from a tailscaled
// config file name previously generated by TailscaledConfigFileNameForCap.
func CapVerFromFileName(name string) (tailcfg.CapabilityVersion, error) {
	if name == "tailscaled" {
		return 0, nil
	}
	var cap tailcfg.CapabilityVersion
	_, err := fmt.Sscanf(name, "cap-%d.hujson", &cap)
	return cap, err
}