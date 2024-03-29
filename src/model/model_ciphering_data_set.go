/*
 * LMF Broadcast
 *
 * LMF Broadcast Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// CipheringDataSet - Represents a Ciphering Data Set.
type CipheringDataSet struct {

	// Ciphering Data Set Identifier.
	CipheringSetID int32 `json:"cipheringSetID"`

	// Ciphering Key.
	CipheringKey string `json:"cipheringKey"`

	// First component of the initial ciphering counter.
	C0 string `json:"c0"`

	// string with format 'bytes' as defined in OpenAPI
	LtePosSibTypes string `json:"ltePosSibTypes,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	NrPosSibTypes string `json:"nrPosSibTypes,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityStartTime time.Time `json:"validityStartTime"`

	// Validity Duration of the Ciphering Data Set.
	ValidityDuration int32 `json:"validityDuration"`

	// string with format 'bytes' as defined in OpenAPI
	TaiList string `json:"taiList,omitempty"`
}
