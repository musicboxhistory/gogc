/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// N2InterfaceAmfInfo - AMF N2 interface information
type N2InterfaceAmfInfo struct {
	Ipv4EndpointAddress []string `json:"ipv4EndpointAddress,omitempty"`

	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`

	// Fully Qualified Domain Name
	AmfName string `json:"amfName,omitempty"`
}
