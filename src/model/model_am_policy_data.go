/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AmPolicyData - Contains the AM policy data for a given subscriber.
type AmPolicyData struct {

	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraInfos map[string]PresenceInfo `json:"praInfos,omitempty"`

	SubscCats []string `json:"subscCats,omitempty"`
}
