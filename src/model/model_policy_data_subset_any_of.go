/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type PolicyDataSubsetAnyOf string

// List of PolicyDataSubsetAnyOf
const (
	AM_POLICY_DATA         PolicyDataSubsetAnyOf = "AM_POLICY_DATA"
	SM_POLICY_DATA         PolicyDataSubsetAnyOf = "SM_POLICY_DATA"
	UE_POLICY_DATA         PolicyDataSubsetAnyOf = "UE_POLICY_DATA"
	UM_DATA                PolicyDataSubsetAnyOf = "UM_DATA"
	OPERATOR_SPECIFIC_DATA PolicyDataSubsetAnyOf = "OPERATOR_SPECIFIC_DATA"
)
