/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// Ipv6PrefixRange - Range of IPv6 prefixes
type Ipv6PrefixRange struct {

	Start Ipv6Prefix `json:"start,omitempty"`

	End Ipv6Prefix `json:"end,omitempty"`
}