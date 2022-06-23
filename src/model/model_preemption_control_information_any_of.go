/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type PreemptionControlInformationAnyOf string

// List of PreemptionControlInformationAnyOf
const (
	MOST_RECENT PreemptionControlInformationAnyOf = "MOST_RECENT"
	LEAST_RECENT PreemptionControlInformationAnyOf = "LEAST_RECENT"
	HIGHEST_BW PreemptionControlInformationAnyOf = "HIGHEST_BW"
)