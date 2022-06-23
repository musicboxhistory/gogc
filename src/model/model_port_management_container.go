/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PortManagementContainer - Contains the port management information container for a port.
type PortManagementContainer struct {

	// string with format 'bytes' as defined in OpenAPI
	PortManCont string `json:"portManCont"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNum int32 `json:"portNum"`
}
