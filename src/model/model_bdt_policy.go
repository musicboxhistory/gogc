/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// BdtPolicy - Represents an Individual BDT policy resource.
type BdtPolicy struct {
	BdtPolData BdtPolicyData `json:"bdtPolData,omitempty"`

	BdtReqData BdtReqData `json:"bdtReqData,omitempty"`
}
