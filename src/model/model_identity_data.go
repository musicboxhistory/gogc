/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// IdentityData - Identity data corresponds to the provided ueId.
type IdentityData struct {
	SupiList []string `json:"supiList,omitempty"`

	GpsiList []string `json:"gpsiList,omitempty"`

	AllowedAfIds []string `json:"allowedAfIds,omitempty"`

	// A map (list of key-value pairs where AppPortId serves as key) of GPSIs.
	ApplicationPortIds map[string]string `json:"applicationPortIds,omitempty"`
}
