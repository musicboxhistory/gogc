/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PreferredSearch - Contains information on whether the returned NFProfiles match the preferred query parameters
type PreferredSearch struct {
	PreferredTaiMatchInd bool `json:"preferredTaiMatchInd,omitempty"`

	PreferredFullPlmnMatchInd bool `json:"preferredFullPlmnMatchInd,omitempty"`

	PreferredApiVersionsMatchInd bool `json:"preferredApiVersionsMatchInd,omitempty"`

	OtherApiVersionsInd bool `json:"otherApiVersionsInd,omitempty"`

	PreferredLocalityMatchInd bool `json:"preferredLocalityMatchInd,omitempty"`

	OtherLocalityInd bool `json:"otherLocalityInd,omitempty"`

	PreferredVendorSpecificFeaturesInd bool `json:"preferredVendorSpecificFeaturesInd,omitempty"`

	PreferredCollocatedNfTypeInd bool `json:"preferredCollocatedNfTypeInd,omitempty"`

	PreferredPgwMatchInd bool `json:"preferredPgwMatchInd,omitempty"`
}
