/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmPolicySnssaiData - Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiData struct {

	Snssai Snssai `json:"snssai"`

	// Session Management Policy data per DNN for all the DNNs of the indicated S-NSSAI. The key of the map is the DNN. 
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty"`

	UeSliceMbr SliceMbr1 `json:"ueSliceMbr,omitempty"`
}