/*
 * Nudm_SSAU
 *
 * Nudm Service Specific Authorization Service.   Â© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type AuthUpdateNotification struct {
	ServiceType ServiceType `json:"serviceType"`

	Snssai Snssai `json:"snssai,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	AuthUpdateInfoList []AuthUpdateInfo `json:"authUpdateInfoList"`
}
