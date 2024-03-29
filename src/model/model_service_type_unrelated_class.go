/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type ServiceTypeUnrelatedClass struct {

	// LCS service type.
	ServiceType int32 `json:"serviceType"`

	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`

	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`

	CodeWordInd CodeWordInd `json:"codeWordInd,omitempty"`

	ValidTimePeriod ValidTimePeriod `json:"validTimePeriod,omitempty"`

	CodeWordList []string `json:"codeWordList,omitempty"`
}
