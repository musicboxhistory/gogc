/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// Notification - Describes a BDT notification.
type Notification struct {

	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`

	// Contains a list of the candidate transfer policies from which the AF may select a new transfer policy due to a network performance is below the criteria set by the operator.
	CandPolicies []TransferPolicy `json:"candPolicies,omitempty"`

	NwAreaInfo NetworkAreaInfo `json:"nwAreaInfo,omitempty"`

	TimeWindow TimeWindow `json:"timeWindow,omitempty"`
}
