/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AccNetChargingAddress - Describes the network entity within the access network performing charging
type AccNetChargingAddress struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AnChargIpv4Addr string `json:"anChargIpv4Addr,omitempty"`

	AnChargIpv6Addr Ipv6Addr `json:"anChargIpv6Addr,omitempty"`
}
