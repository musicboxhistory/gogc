/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AdditionalAccessInfo - Indicates the combination of additional Access Type and RAT Type for a MA PDU session.
type AdditionalAccessInfo struct {
	AccessType AccessType `json:"accessType"`

	RatType RatType `json:"ratType,omitempty"`
}
