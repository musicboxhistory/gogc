/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SequenceNumber - Contains the SQN.
type SequenceNumber struct {

	SqnScheme SqnScheme `json:"sqnScheme,omitempty"`

	Sqn string `json:"sqn,omitempty"`

	LastIndexes map[string]int32 `json:"lastIndexes,omitempty"`

	IndLength int32 `json:"indLength,omitempty"`

	DifSign Sign `json:"difSign,omitempty"`
}
