/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type DnnInfo struct {

	Dnn AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"dnn"`

	DefaultDnnIndicator bool `json:"defaultDnnIndicator,omitempty"`

	LboRoamingAllowed bool `json:"lboRoamingAllowed,omitempty"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	DnnBarred bool `json:"dnnBarred,omitempty"`

	InvokeNefInd bool `json:"invokeNefInd,omitempty"`

	SmfList []string `json:"smfList,omitempty"`

	SameSmfInd bool `json:"sameSmfInd,omitempty"`
}
