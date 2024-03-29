/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmContextUpdatedData - Data within Update SM Context Response
type SmContextUpdatedData struct {

	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	HoState HoState `json:"hoState,omitempty"`

	ReleaseEbiList []int32 `json:"releaseEbiList,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	EpsBearerSetup []string `json:"epsBearerSetup,omitempty"`

	DataForwarding bool `json:"dataForwarding,omitempty"`

	N3DlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`

	N3UlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`

	Cause Cause `json:"cause,omitempty"`

	MaAcceptedInd bool `json:"maAcceptedInd,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	ForwardingFTeid string `json:"forwardingFTeid,omitempty"`

	ForwardingBearerContexts []string `json:"forwardingBearerContexts,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SelectedSmfId string `json:"selectedSmfId,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SelectedOldSmfId string `json:"selectedOldSmfId,omitempty"`

	// String providing an URI formatted according to RFC 3986
	InterPlmnApiRoot string `json:"interPlmnApiRoot,omitempty"`
}
