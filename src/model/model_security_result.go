/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SecurityResult - Security Result
type SecurityResult struct {

	IntegrityProtectionResult ProtectionResult `json:"integrityProtectionResult,omitempty"`

	ConfidentialityProtectionResult ProtectionResult `json:"confidentialityProtectionResult,omitempty"`
}