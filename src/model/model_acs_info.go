/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AcsInfo - The ACS information for the 5G-RG is defined in BBF TR-069 [42] or in BBF TR-369
type AcsInfo struct {

	// String providing an URI formatted according to RFC 3986
	AcsUrl string `json:"acsUrl,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AcsIpv4Addr string `json:"acsIpv4Addr,omitempty"`

	AcsIpv6Addr Ipv6Addr `json:"acsIpv6Addr,omitempty"`
}
