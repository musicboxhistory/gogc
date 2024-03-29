/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// Flows - Identifies the flows.
type Flows struct {
	ContVers []int32 `json:"contVers,omitempty"`

	FNums []int32 `json:"fNums,omitempty"`

	MedCompN int32 `json:"medCompN"`
}
