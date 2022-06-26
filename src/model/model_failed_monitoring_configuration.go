/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// FailedMonitoringConfiguration - Contains the event type and failed cause of the failed Monitoring Configuration in the EE subscription
type FailedMonitoringConfiguration struct {
	EventType EventType `json:"eventType"`

	FailedCause FailedCause `json:"failedCause"`
}
