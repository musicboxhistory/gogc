/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type CreatedEeSubscription struct {

	EeSubscription EeSubscription `json:"eeSubscription"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumberOfUes int32 `json:"numberOfUes,omitempty"`

	EventReports []MonitoringReport `json:"eventReports,omitempty"`

	EpcStatusInd bool `json:"epcStatusInd,omitempty"`

	// A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration
	FailedMonitoringConfigs map[string]FailedMonitoringConfiguration `json:"failedMonitoringConfigs,omitempty"`

	// A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration, the key value \"ALL\" may be used to identify a map entry which contains the failed cause of the EE subscription was not successful in EPC domain.
	FailedMoniConfigsEPC map[string]FailedMonitoringConfiguration `json:"failedMoniConfigsEPC,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`
}
