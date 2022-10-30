/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type ReportedEventTypeAnyOf string

// List of ReportedEventTypeAnyOf
const (
	PERIODIC_EVENT ReportedEventTypeAnyOf = "PERIODIC_EVENT"
	ENTERING_AREA_EVENT ReportedEventTypeAnyOf = "ENTERING_AREA_EVENT"
	LEAVING_AREA_EVENT ReportedEventTypeAnyOf = "LEAVING_AREA_EVENT"
	BEING_INSIDE_AREA_EVENT ReportedEventTypeAnyOf = "BEING_INSIDE_AREA_EVENT"
	MOTION_EVENT ReportedEventTypeAnyOf = "MOTION_EVENT"
	MAXIMUM_INTERVAL_EXPIRATION_EVENT ReportedEventTypeAnyOf = "MAXIMUM_INTERVAL_EXPIRATION_EVENT"
	LOCATION_CANCELLATION_EVENT ReportedEventTypeAnyOf = "LOCATION_CANCELLATION_EVENT"
)