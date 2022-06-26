/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// TrafficControlData - Contains parameters determining how flows associated with a PCC Rule are treated (e.g. blocked, redirected, etc).
type TrafficControlData struct {

	// Univocally identifies the traffic control policy data within a PDU session.
	TcId string `json:"tcId"`

	FlowStatus FlowStatus `json:"flowStatus,omitempty"`

	RedirectInfo RedirectInformation `json:"redirectInfo,omitempty"`

	AddRedirectInfo []RedirectInformation `json:"addRedirectInfo,omitempty"`

	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `json:"muteNotif,omitempty"`

	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl *string `json:"trafficSteeringPolIdDl,omitempty"`

	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl *string `json:"trafficSteeringPolIdUl,omitempty"`

	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs *[]RouteToLocation `json:"routeToLocs,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	MaxAllowedUpLat *int32 `json:"maxAllowedUpLat,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos *[]EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`

	TraffCorreInd bool `json:"traffCorreInd,omitempty"`

	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty"`

	// indicating a time in seconds.
	SimConnTerm int32 `json:"simConnTerm,omitempty"`

	UpPathChgEvent *UpPathChgEvent `json:"upPathChgEvent,omitempty"`

	SteerFun SteeringFunctionality `json:"steerFun,omitempty"`

	SteerModeDl SteeringMode `json:"steerModeDl,omitempty"`

	SteerModeUl SteeringMode `json:"steerModeUl,omitempty"`

	MulAccCtrl MulticastAccessControl `json:"mulAccCtrl,omitempty"`
}
