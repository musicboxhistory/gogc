/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// UeLocationInfo - Indicates location information of a UE.
type UeLocationInfo struct {

	LocationEstimate GeographicArea `json:"locationEstimate,omitempty"`

	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate int32 `json:"ageOfLocationEstimate,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate time.Time `json:"timestampOfLocationEstimate,omitempty"`

	VelocityEstimate VelocityEstimate `json:"velocityEstimate,omitempty"`

	// Indicates value of the age of the location estimate.
	AgeOfVelocityEstimate int32 `json:"ageOfVelocityEstimate,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfVelocityEstimate time.Time `json:"timestampOfVelocityEstimate,omitempty"`
}
