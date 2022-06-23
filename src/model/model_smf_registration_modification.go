/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmfRegistrationModification - Contains attributes of SmfRegistration that can be modified using PATCH
type SmfRegistrationModification struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmfInstanceId string `json:"smfInstanceId"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"  set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or \"set<SetID>.  <NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with <MCC> encoded as defined in clause 5.4.2  (\"Mcc\" data type definition) <MNC> encoded as defined in clause 5.4.2 (\"Mnc\" data type  definition) <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but  with lower case characters <Set ID> encoded as a string of characters consisting of alphabetic  characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an  alphabetic character or a digit.  
	SmfSetId string `json:"smfSetId,omitempty"`

	PgwFqdn FqdnRm `json:"pgwFqdn,omitempty"`
}
