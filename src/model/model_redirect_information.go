/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RedirectInformation - Contains the redirect information.
type RedirectInformation struct {

	// Indicates the redirect is enable.
	RedirectEnabled bool `json:"redirectEnabled,omitempty"`

	RedirectAddressType RedirectAddressType `json:"redirectAddressType,omitempty"`

	// Indicates the address of the redirect server. If \"redirectAddressType\" attribute indicates the IPV4_ADDR, the encoding is the same as the Ipv4Addr data type defined in 3GPP TS 29.571.If \"redirectAddressType\" attribute indicates the IPV6_ADDR, the encoding is the same as the Ipv6Addr data type defined in 3GPP TS 29.571.If \"redirectAddressType\" attribute indicates the URL or SIP_URI, the encoding is the same as the Uri data type defined in 3GPP TS 29.571.
	RedirectServerAddress string `json:"redirectServerAddress,omitempty"`
}
