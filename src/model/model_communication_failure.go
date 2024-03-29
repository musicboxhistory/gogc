/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// CommunicationFailure - Describes a communication failure detected by AMF
type CommunicationFailure struct {
	NasReleaseCode string `json:"nasReleaseCode,omitempty"`

	RanReleaseCode NgApCause `json:"ranReleaseCode,omitempty"`
}
