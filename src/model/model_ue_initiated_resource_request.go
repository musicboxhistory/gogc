/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UeInitiatedResourceRequest - Indicates that a UE requests specific QoS handling for the selected SDF.
type UeInitiatedResourceRequest struct {

	PccRuleId string `json:"pccRuleId,omitempty"`

	RuleOp RuleOperation `json:"ruleOp"`

	Precedence int32 `json:"precedence,omitempty"`

	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`

	ReqQos RequestedQos `json:"reqQos,omitempty"`
}