/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// EventsNotification - Describes the notification of a matched event.
type EventsNotification struct {

	// Includes the detected application report.
	AdReports []AppDetectionReport `json:"adReports,omitempty"`

	AccessType AccessType `json:"accessType,omitempty"`

	AddAccessInfo AdditionalAccessInfo `json:"addAccessInfo,omitempty"`

	RelAccessInfo AdditionalAccessInfo `json:"relAccessInfo,omitempty"`

	AnChargAddr AccNetChargingAddress `json:"anChargAddr,omitempty"`

	AnChargIds []AccessNetChargingIdentifier `json:"anChargIds,omitempty"`

	AnGwAddr AnGwAddress `json:"anGwAddr,omitempty"`

	// String providing an URI formatted according to RFC 3986
	EvSubsUri string `json:"evSubsUri"`

	EvNotifs []AfEventNotification `json:"evNotifs"`

	FailedResourcAllocReports []ResourcesAllocationInfo `json:"failedResourcAllocReports,omitempty"`

	SuccResourcAllocReports []ResourcesAllocationInfo `json:"succResourcAllocReports,omitempty"`

	NoNetLocSupp NetLocAccessSupport `json:"noNetLocSupp,omitempty"`

	OutOfCredReports []OutOfCreditInformation `json:"outOfCredReports,omitempty"`

	PlmnId PlmnIdNid `json:"plmnId,omitempty"`

	QncReports []QosNotificationControlInfo `json:"qncReports,omitempty"`

	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`

	UeLoc UserLocation `json:"ueLoc,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UeLocTime time.Time `json:"ueLocTime,omitempty"`

	// String with format \"<time-numoffset>\" optionally appended by \"<daylightSavingTime>\", where  -  <time-numoffset> shall represent the time zone adjusted for daylight saving time and be encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - <daylightSavingTime> shall represent the adjustment that has been made and shall be encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.  The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone string `json:"ueTimeZone,omitempty"`

	UsgRep AccumulatedUsage `json:"usgRep,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}
