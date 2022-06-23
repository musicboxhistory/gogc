/*
 * Nudm_NIDDAU
 *
 * Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0-alpha.3
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// NiddAuthUpdateInfo - Represents NIDD authorization update information.
type NiddAuthUpdateInfo struct {

	AuthorizationData AuthorizationData `json:"authorizationData"`

	InvalidityInd bool `json:"invalidityInd,omitempty"`
}
