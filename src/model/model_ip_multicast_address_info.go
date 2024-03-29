/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// IpMulticastAddressInfo - Contains the IP multicast addressing information.
type IpMulticastAddressInfo struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SrcIpv4Addr string `json:"srcIpv4Addr,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4MulAddr string `json:"ipv4MulAddr,omitempty"`

	SrcIpv6Addr Ipv6Addr `json:"srcIpv6Addr,omitempty"`

	Ipv6MulAddr Ipv6Addr `json:"ipv6MulAddr,omitempty"`
}
