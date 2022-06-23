/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AppAmContextUpdateData - Describes the modifications to an Individual Application AM resource.
type AppAmContextUpdateData struct {

	// String providing an URI formatted according to RFC 3986
	TermNotifUri string `json:"termNotifUri,omitempty"`

	EvSubsc *AmEventsSubscDataRm `json:"evSubsc,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	Expiry *int32 `json:"expiry,omitempty"`

	// Indicates whether high throughput is desired for the indicated UE traffic.
	HighThruInd *bool `json:"highThruInd,omitempty"`

	// Identifies a list of Tracking Areas per serving network where service is allowed.
	CovReq *[]ServiceAreaCoverageInfo `json:"covReq,omitempty"`

	AsTimeDisParam *AsTimeDistributionParam `json:"asTimeDisParam,omitempty"`
}
