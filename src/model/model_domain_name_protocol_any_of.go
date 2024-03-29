/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type DomainNameProtocolAnyOf string

// List of DomainNameProtocolAnyOf
const (
	DNS_QNAME DomainNameProtocolAnyOf = "DNS_QNAME"
	TLS_SNI   DomainNameProtocolAnyOf = "TLS_SNI"
	TLS_SAN   DomainNameProtocolAnyOf = "TLS_SAN"
	TSL_SCN   DomainNameProtocolAnyOf = "TSL_SCN"
)
