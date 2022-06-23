/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type DnnSelectionModeAnyOf string

// List of DnnSelectionModeAnyOf
const (
	VERIFIED DnnSelectionModeAnyOf = "VERIFIED"
	UE_DNN_NOT_VERIFIED DnnSelectionModeAnyOf = "UE_DNN_NOT_VERIFIED"
	NW_DNN_NOT_VERIFIED DnnSelectionModeAnyOf = "NW_DNN_NOT_VERIFIED"
)
