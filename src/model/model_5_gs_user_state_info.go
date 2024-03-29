/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// Model5GsUserStateInfo - Represents the 5GS User state of the UE for an access type
type Model5GsUserStateInfo struct {
	Var5gsUserState Model5GsUserState `json:"5gsUserState"`

	AccessType AccessType `json:"accessType"`
}
