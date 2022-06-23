/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PartialSuccessReport - Includes the information reported by the SMF when some of the PCC rules and/or session rules are not successfully installed/activated.
type PartialSuccessReport struct {

	FailureCause FailureCause `json:"failureCause"`

	// Information about the PCC rules provisioned by the PCF not successfully installed/activated.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`

	// Information about the session rules provisioned by the PCF not successfully installed.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`

	UeCampingRep UeCampingRep `json:"ueCampingRep,omitempty"`

	// Contains the type(s) of failed policy decision and/or condition data.
	PolicyDecFailureReports []PolicyDecisionFailureCode `json:"policyDecFailureReports,omitempty"`

	// Indicates the invalid parameters for the reported type(s) of the failed policy decision and/or condition data.
	InvalidPolicyDecs []InvalidParam `json:"invalidPolicyDecs,omitempty"`
}
