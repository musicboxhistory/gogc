/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SpatialValidity - Describes explicitly the route to an Application location.
type SpatialValidity struct {

	// Defines the presence information provisioned by the AF. The praId attribute within the PresenceInfo data type is the key of the map.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList"`
}