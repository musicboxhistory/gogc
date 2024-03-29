/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DispersionArea - Dispersion Area
type DispersionArea struct {
	TaiList []Tai `json:"taiList,omitempty"`

	NcgiList []Ncgi `json:"ncgiList,omitempty"`

	EcgiList []Ecgi `json:"ecgiList,omitempty"`

	N3gaInd bool `json:"n3gaInd,omitempty"`
}
