/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// VsmfUpdateData - Data within Update Request towards V-SMF, or from SMF to I-SMF
type VsmfUpdateData struct {

	RequestIndication RequestIndication `json:"requestIndication"`

	SessionAmbr Ambr `json:"sessionAmbr,omitempty"`

	QosFlowsAddModRequestList []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`

	QosFlowsRelRequestList []QosFlowReleaseRequestItem `json:"qosFlowsRelRequestList,omitempty"`

	EpsBearerInfo []EpsBearerInfo `json:"epsBearerInfo,omitempty"`

	AssignEbiList []Arp `json:"assignEbiList,omitempty"`

	RevokeEbiList []int32 `json:"revokeEbiList,omitempty"`

	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`

	// Procedure Transaction Identifier
	Pti int32 `json:"pti,omitempty"`

	N1SmInfoToUe RefToBinaryData `json:"n1SmInfoToUe,omitempty"`

	AlwaysOnGranted bool `json:"alwaysOnGranted,omitempty"`

	// String providing an URI formatted according to RFC 3986
	HsmfPduSessionUri string `json:"hsmfPduSessionUri,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	Cause Cause `json:"cause,omitempty"`

	N1smCause string `json:"n1smCause,omitempty"`

	// indicating a time in seconds.
	BackOffTimer int32 `json:"backOffTimer,omitempty"`

	MaReleaseInd MaReleaseIndication `json:"maReleaseInd,omitempty"`

	MaAcceptedInd bool `json:"maAcceptedInd,omitempty"`

	AdditionalCnTunnelInfo TunnelInfo `json:"additionalCnTunnelInfo,omitempty"`

	DnaiList []string `json:"dnaiList,omitempty"`

	N4Info N4Information `json:"n4Info,omitempty"`

	N4InfoExt1 N4Information `json:"n4InfoExt1,omitempty"`

	N4InfoExt2 N4Information `json:"n4InfoExt2,omitempty"`

	N4InfoExt3 N4Information `json:"n4InfoExt3,omitempty"`

	SmallDataRateControlEnabled bool `json:"smallDataRateControlEnabled,omitempty"`

	QosMonitoringInfo QosMonitoringInfo `json:"qosMonitoringInfo,omitempty"`

	EpsPdnCnxInfo EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`
}
