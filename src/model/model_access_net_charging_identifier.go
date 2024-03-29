/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AccessNetChargingIdentifier - Describes the access network charging identifier.
type AccessNetChargingIdentifier struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	AccNetChaIdValue int32 `json:"accNetChaIdValue,omitempty"`

	// A character string containing the access network charging identifier.
	AccNetChargIdString string `json:"accNetChargIdString,omitempty"`

	Flows []Flows `json:"flows,omitempty"`
}
