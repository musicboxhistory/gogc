/*
 * Naf_EventExposure
 *
 * AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type MetricsReportingConfiguration struct {

	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	MetricsReportingConfigurationId string `json:"metricsReportingConfigurationId"`

	// String providing an URI formatted according to RFC 3986
	Scheme string `json:"scheme"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	DataNetworkName string `json:"dataNetworkName,omitempty"`

	// indicating a time in seconds.
	ReportingInterval int32 `json:"reportingInterval,omitempty"`

	SamplePercentage float32 `json:"samplePercentage,omitempty"`

	UrlFilters []string `json:"urlFilters,omitempty"`

	Metrics []string `json:"metrics,omitempty"`
}
