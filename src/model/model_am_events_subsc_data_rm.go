/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AmEventsSubscDataRm - This data type is defined in the same way as the AmEventsSubscData but with the OpenAPI nullable property set to true.
type AmEventsSubscDataRm struct {

	// String providing an URI formatted according to RFC 3986
	EventNotifUri string `json:"eventNotifUri,omitempty"`

	Events []AmEventData `json:"events,omitempty"`
}
