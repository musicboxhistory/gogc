/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmContextRetrievedData - Data within Retrieve SM Context Response
type SmContextRetrievedData struct {

	// UE EPS PDN Connection container from SMF to AMF
	UeEpsPdnConnection string `json:"ueEpsPdnConnection"`

	SmContext SmContext `json:"smContext,omitempty"`

	SmallDataRateStatus SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`

	ApnRateStatus ApnRateStatus `json:"apnRateStatus,omitempty"`

	DlDataWaitingInd bool `json:"dlDataWaitingInd,omitempty"`

	AfCoordinationInfo AfCoordinationInfo `json:"afCoordinationInfo,omitempty"`
}
