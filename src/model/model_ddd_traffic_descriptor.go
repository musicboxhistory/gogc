/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DddTrafficDescriptor - Contains a Traffic Descriptor.
type DddTrafficDescriptor struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	MacAddr string `json:"macAddr,omitempty"`
}
