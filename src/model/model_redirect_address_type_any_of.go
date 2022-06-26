/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type RedirectAddressTypeAnyOf string

// List of RedirectAddressTypeAnyOf
const (
	IPV4_ADDR RedirectAddressTypeAnyOf = "IPV4_ADDR"
	IPV6_ADDR RedirectAddressTypeAnyOf = "IPV6_ADDR"
	URL       RedirectAddressTypeAnyOf = "URL"
	SIP_URI   RedirectAddressTypeAnyOf = "SIP_URI"
)
