/*
 * Naf_EventExposure
 *
 * AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// MsAccessActivityCollection - Contains Media Streaming access activity collected for an UE Application via AF.
type MsAccessActivityCollection struct {
	MsAccActs []string `json:"msAccActs"`
}
