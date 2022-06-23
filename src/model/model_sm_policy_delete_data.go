/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// SmPolicyDeleteData - Contains the parameters to be sent to the PCF when an individual SM policy is deleted.
type SmPolicyDeleteData struct {

	UserLocationInfo UserLocation `json:"userLocationInfo,omitempty"`

	// String with format \"<time-numoffset>\" optionally appended by \"<daylightSavingTime>\", where  -  <time-numoffset> shall represent the time zone adjusted for daylight saving time and be encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - <daylightSavingTime> shall represent the adjustment that has been made and shall be encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.  The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone string `json:"ueTimeZone,omitempty"`

	ServingNetwork PlmnIdNid `json:"servingNetwork,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UserLocationInfoTime time.Time `json:"userLocationInfoTime,omitempty"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`

	PduSessRelCause PduSessionRelCause `json:"pduSessRelCause,omitempty"`

	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`
}
