/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AppAmContextRespData - It represents a response to a modification or creation request of an Individual Application AM resource. It may contain the notification of the already met events.
type AppAmContextRespData struct {

	// Contains the AM Policy Events Subscription resource identifier related to the event notification.
	AppAmContextId string `json:"appAmContextId,omitempty"`

	RepEvents []AmEventNotification `json:"repEvents"`
}
