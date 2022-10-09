/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmContextReleaseData - Data within Release SM Context Request
type SmContextReleaseData struct {

	Cause Cause `json:"cause,omitempty"`

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue int32 `json:"5gMmCauseValue,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`

	// String with format \"<time-numoffset>\" optionally appended by \"<daylightSavingTime>\", where  -  <time-numoffset> shall represent the time zone adjusted for daylight saving time and be encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - <daylightSavingTime> shall represent the adjustment that has been made and shall be encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.  The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation UserLocation `json:"addUeLocation,omitempty"`

	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	IsmfReleaseOnly bool `json:"ismfReleaseOnly,omitempty"`
}
