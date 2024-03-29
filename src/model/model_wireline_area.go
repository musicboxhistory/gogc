/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// WirelineArea - One and only one of the \"globLineIds\", \"hfcNIds\", \"areaCodeB\" and \"areaCodeC\" attributes shall be included in a WirelineArea data structure
type WirelineArea struct {
	GlobalLineIds []string `json:"globalLineIds,omitempty"`

	HfcNIds []string `json:"hfcNIds,omitempty"`

	// Values are operator specific.
	AreaCodeB string `json:"areaCodeB,omitempty"`

	// Values are operator specific.
	AreaCodeC string `json:"areaCodeC,omitempty"`
}
