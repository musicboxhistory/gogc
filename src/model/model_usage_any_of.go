/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type UsageAnyOf string

// List of UsageAnyOf
const (
	UNSUCCESS UsageAnyOf = "UNSUCCESS"
	SUCCESS_RESULTS_NOT_USED UsageAnyOf = "SUCCESS_RESULTS_NOT_USED"
	SUCCESS_RESULTS_USED_TO_VERIFY_LOCATION UsageAnyOf = "SUCCESS_RESULTS_USED_TO_VERIFY_LOCATION"
	SUCCESS_RESULTS_USED_TO_GENERATE_LOCATION UsageAnyOf = "SUCCESS_RESULTS_USED_TO_GENERATE_LOCATION"
	SUCCESS_METHOD_NOT_DETERMINED UsageAnyOf = "SUCCESS_METHOD_NOT_DETERMINED"
)
