/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// QosFlowItem - Individual QoS flow
type QosFlowItem struct {

	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi"`

	Cause Cause `json:"cause,omitempty"`

	CurrentQosProfileIndex int32 `json:"currentQosProfileIndex,omitempty"`

	NullQoSProfileIndex bool `json:"nullQoSProfileIndex,omitempty"`

	NgApCause NgApCause `json:"ngApCause,omitempty"`
}
