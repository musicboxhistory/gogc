/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RuleOperation - Possible values are - CREATE_PCC_RULE: Indicates to create a new PCC rule to reserve the resource requested by the UE.  - DELETE_PCC_RULE: Indicates to delete a PCC rule corresponding to reserve the resource requested by the UE. - MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS: Indicates to modify the PCC rule by adding new packet filter(s). - MODIFY_ PCC_RULE_AND_REPLACE_PACKET_FILTERS: Indicates to modify the PCC rule by replacing the existing packet filter(s). - MODIFY_ PCC_RULE_AND_DELETE_PACKET_FILTERS: Indicates to modify the PCC rule by deleting the existing packet filter(s). - MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS: Indicates to modify the PCC rule by modifying the QoS of the PCC rule.
type RuleOperation struct {
}
