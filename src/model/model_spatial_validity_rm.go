/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SpatialValidityRm - This data type is defined in the same way as the SpatialValidity data type, but with the OpenAPI nullable property set to true.
type SpatialValidityRm struct {

	// Defines the presence information provisioned by the AF. The praId attribute within the PresenceInfo data type is the key of the map.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList"`
}
