/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// InterFreqTargetInfo - Indicates the Inter Frequency Target information.
type InterFreqTargetInfo struct {

	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331. 
	DlCarrierFreq int32 `json:"dlCarrierFreq"`

	// When present, this IE shall contain a list of the physical cell identities where the UE is requested to perform measurement logging for the indicated frequency.
	CellIdList []int32 `json:"cellIdList,omitempty"`
}
