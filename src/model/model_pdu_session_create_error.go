/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// PduSessionCreateError - Error within Create Response
type PduSessionCreateError struct {

	Error ProblemDetails `json:"error"`

	N1smCause string `json:"n1smCause,omitempty"`

	N1SmInfoToUe RefToBinaryData `json:"n1SmInfoToUe,omitempty"`

	// indicating a time in seconds.
	BackOffTimer int32 `json:"backOffTimer,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime time.Time `json:"recoveryTime,omitempty"`
}
