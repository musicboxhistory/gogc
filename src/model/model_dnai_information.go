/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DnaiInformation - DNAI Information
type DnaiInformation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	NoDnaiChangeInd bool `json:"noDnaiChangeInd,omitempty"`

	NoLocalPsaChangeInd bool `json:"noLocalPsaChangeInd,omitempty"`
}
