/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ScheduledCommunicationTimeRm - This data type is defined in the same way as the 'ScheduledCommunicationTime' data type, but with the OpenAPI 'nullable: true' property.
type ScheduledCommunicationTimeRm struct {

	// Identifies the day(s) of the week. If absent, it indicates every day of the week.
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty"`

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayEnd string `json:"timeOfDayEnd,omitempty"`
}
