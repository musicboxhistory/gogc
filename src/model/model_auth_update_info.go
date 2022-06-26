/*
 * Nudm_SSAU
 *
 * Nudm Service Specific Authorization Service.   Â© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type AuthUpdateInfo struct {
	AuthorizationData AuthorizationData `json:"authorizationData"`

	InvalidityInd bool `json:"invalidityInd,omitempty"`
}
