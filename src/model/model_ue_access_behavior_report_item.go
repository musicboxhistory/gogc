/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UeAccessBehaviorReportItem - Report Item for UE Access Behavior Trends event.
type UeAccessBehaviorReportItem struct {
	StateTransitionType AccessStateTransitionType `json:"stateTransitionType"`

	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`

	// indicating a time in seconds.
	Duration int32 `json:"duration"`
}
