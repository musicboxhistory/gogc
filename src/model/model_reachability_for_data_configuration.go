/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type ReachabilityForDataConfiguration struct {
	ReportCfg ReachabilityForDataReportConfig `json:"reportCfg"`

	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty"`
}
