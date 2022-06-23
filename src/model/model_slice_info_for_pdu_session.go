/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SliceInfoForPduSession - Contains the slice information requested during PDU Session establishment procedure
type SliceInfoForPduSession struct {

	SNssai Snssai `json:"sNssai"`

	RoamingIndication RoamingIndication `json:"roamingIndication"`

	HomeSnssai Snssai `json:"homeSnssai,omitempty"`
}