/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RoamingRestrictions - Indicates if access is allowed to a given serving network, e.g. a PLMN (MCC, MNC) or an SNPN (MCC, MNC, NID).
type RoamingRestrictions struct {
	AccessAllowed bool `json:"accessAllowed,omitempty"`
}
