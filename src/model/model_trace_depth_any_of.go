/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type TraceDepthAnyOf string

// List of TraceDepthAnyOf
const (
	MINIMUM TraceDepthAnyOf = "MINIMUM"
	MEDIUM TraceDepthAnyOf = "MEDIUM"
	MAXIMUM TraceDepthAnyOf = "MAXIMUM"
	MINIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MINIMUM_WO_VENDOR_EXTENSION"
	MEDIUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MEDIUM_WO_VENDOR_EXTENSION"
	MAXIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MAXIMUM_WO_VENDOR_EXTENSION"
)