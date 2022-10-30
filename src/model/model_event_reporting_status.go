/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// EventReportingStatus - Indicates the status of event reporting.
type EventReportingStatus struct {

	// Number of event reports received from the target UE.
	EventReportCounter int32 `json:"eventReportCounter,omitempty"`

	// Duration of event reporting.
	EventReportDuration int32 `json:"eventReportDuration,omitempty"`
}