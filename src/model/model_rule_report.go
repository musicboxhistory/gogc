/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RuleReport - Reports the status of PCC.
type RuleReport struct {

	// Contains the identifier of the affected PCC rule(s).
	PccRuleIds []string `json:"pccRuleIds"`

	RuleStatus RuleStatus `json:"ruleStatus"`

	// Indicates the version of a PCC rule.
	ContVers []int32 `json:"contVers,omitempty"`

	FailureCode FailureCode `json:"failureCode,omitempty"`

	FinUnitAct FinalUnitAction `json:"finUnitAct,omitempty"`

	// indicates the RAN or NAS release cause code information.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	AltQosParamId string `json:"altQosParamId,omitempty"`
}
