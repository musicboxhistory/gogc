/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// BackupAmfInfo - Provides details of the Backup AMF.
type BackupAmfInfo struct {

	// Fully Qualified Domain Name
	BackupAmf string `json:"backupAmf"`

	// If present, this IE shall contain the list of GUAMI(s) (supported by the AMF) for which the backupAmf IE applies.
	GuamiList []Guami `json:"guamiList,omitempty"`
}
