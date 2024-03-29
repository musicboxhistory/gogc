/*
 * Npcf_BDTPolicyControl Service API
 *
 * PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PatchBdtPolicy - Describes the updates in authorization data of an Individual BDT Policy created by the PCF.
type PatchBdtPolicy struct {
	BdtPolData BdtPolicyDataPatch `json:"bdtPolData,omitempty"`

	BdtReqData BdtReqDataPatch `json:"bdtReqData,omitempty"`
}
