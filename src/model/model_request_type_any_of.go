/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type RequestTypeAnyOf string

// List of RequestTypeAnyOf
const (
	INITIAL_REQUEST RequestTypeAnyOf = "INITIAL_REQUEST"
	EXISTING_PDU_SESSION RequestTypeAnyOf = "EXISTING_PDU_SESSION"
	INITIAL_EMERGENCY_REQUEST RequestTypeAnyOf = "INITIAL_EMERGENCY_REQUEST"
	EXISTING_EMERGENCY_PDU_SESSION RequestTypeAnyOf = "EXISTING_EMERGENCY_PDU_SESSION"
)