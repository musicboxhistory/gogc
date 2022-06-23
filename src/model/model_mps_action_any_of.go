/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type MpsActionAnyOf string

// List of MpsActionAnyOf
const (
	DISABLE_MPS_FOR_DTS MpsActionAnyOf = "DISABLE_MPS_FOR_DTS"
	ENABLE_MPS_FOR_DTS MpsActionAnyOf = "ENABLE_MPS_FOR_DTS"
	AUTHORIZE_AND_ENABLE_MPS_FOR_DTS MpsActionAnyOf = "AUTHORIZE_AND_ENABLE_MPS_FOR_DTS"
)