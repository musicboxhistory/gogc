/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// MbsSessionId - MBS Session Identifier
type MbsSessionId struct {

	Tmgi Tmgi `json:"tmgi,omitempty"`

	Ssm Ssm `json:"ssm,omitempty"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty"`
}
