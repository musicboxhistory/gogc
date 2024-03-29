/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type RuleOperationAnyOf string

// List of RuleOperationAnyOf
const (
	CREATE_PCC_RULE                               RuleOperationAnyOf = "CREATE_PCC_RULE"
	DELETE_PCC_RULE                               RuleOperationAnyOf = "DELETE_PCC_RULE"
	MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS        RuleOperationAnyOf = "MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS"
	MODIFY__PCC_RULE_AND_REPLACE_PACKET_FILTERS   RuleOperationAnyOf = "MODIFY_ PCC_RULE_AND_REPLACE_PACKET_FILTERS"
	MODIFY__PCC_RULE_AND_DELETE_PACKET_FILTERS    RuleOperationAnyOf = "MODIFY_ PCC_RULE_AND_DELETE_PACKET_FILTERS"
	MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS RuleOperationAnyOf = "MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS"
)
