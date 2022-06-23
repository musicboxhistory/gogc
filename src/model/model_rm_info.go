/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RmInfo - Represents the registration state of a UE for an access type
type RmInfo struct {

	RmState RmState `json:"rmState"`

	AccessType AccessType `json:"accessType"`
}
