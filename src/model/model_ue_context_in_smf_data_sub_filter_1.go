/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UeContextInSmfDataSubFilter1 - UE Context In Smf Data Subscription Filter.
type UeContextInSmfDataSubFilter1 struct {

	DnnList []string `json:"dnnList,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	EmergencyInd bool `json:"emergencyInd,omitempty"`
}
