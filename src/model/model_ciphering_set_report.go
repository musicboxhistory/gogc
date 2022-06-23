/*
 * LMF Broadcast
 *
 * LMF Broadcast Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// CipheringSetReport - Represents a report of Ciphering Data Set storage.
type CipheringSetReport struct {

	// Ciphering Data Set Identifier.
	CipheringSetID int32 `json:"cipheringSetID"`

	StorageOutcome StorageOutcome `json:"storageOutcome"`
}
