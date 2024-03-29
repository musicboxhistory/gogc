/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// TngfInfo - Infomation of the TNGF endpoints
type TngfInfo struct {

	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`

	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`

	// Fully Qualified Domain Name
	EndpointFqdn string `json:"endpointFqdn,omitempty"`
}
