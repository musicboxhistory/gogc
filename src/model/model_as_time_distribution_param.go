/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AsTimeDistributionParam - Contains the 5G acess stratum time distribution parameters.
type AsTimeDistributionParam struct {
	AsTimeDistInd bool `json:"asTimeDistInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	UuErrorBudget *int32 `json:"uuErrorBudget,omitempty"`
}
