/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type AccessRightStatusAnyOf string

// List of AccessRightStatusAnyOf
const (
	FULLY_ALLOWED   AccessRightStatusAnyOf = "FULLY_ALLOWED"
	PREVIEW_ALLOWED AccessRightStatusAnyOf = "PREVIEW_ALLOWED"
	NO_ALLOWED      AccessRightStatusAnyOf = "NO_ALLOWED"
)
