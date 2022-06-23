/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RequestedRuleData - Contains rule data requested by the PCF to receive information associated with PCC rule(s).
type RequestedRuleData struct {

	// An array of PCC rule id references to the PCC rules associated with the control data.
	RefPccRuleIds []string `json:"refPccRuleIds"`

	// Array of requested rule data type elements indicating what type of rule data is requested for the corresponding referenced PCC rules.
	ReqData []RequestedRuleDataType `json:"reqData"`
}
