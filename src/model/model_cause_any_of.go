/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type CauseAnyOf string

// List of CauseAnyOf
const (
	REL_DUE_TO_HO CauseAnyOf = "REL_DUE_TO_HO"
	EPS_FALLBACK CauseAnyOf = "EPS_FALLBACK"
	REL_DUE_TO_UP_SEC CauseAnyOf = "REL_DUE_TO_UP_SEC"
	DNN_CONGESTION CauseAnyOf = "DNN_CONGESTION"
	S_NSSAI_CONGESTION CauseAnyOf = "S_NSSAI_CONGESTION"
	REL_DUE_TO_REACTIVATION CauseAnyOf = "REL_DUE_TO_REACTIVATION"
	_5_G_AN_NOT_RESPONDING CauseAnyOf = "5G_AN_NOT_RESPONDING"
	REL_DUE_TO_SLICE_NOT_AVAILABLE CauseAnyOf = "REL_DUE_TO_SLICE_NOT_AVAILABLE"
	REL_DUE_TO_DUPLICATE_SESSION_ID CauseAnyOf = "REL_DUE_TO_DUPLICATE_SESSION_ID"
	PDU_SESSION_STATUS_MISMATCH CauseAnyOf = "PDU_SESSION_STATUS_MISMATCH"
	HO_FAILURE CauseAnyOf = "HO_FAILURE"
	INSUFFICIENT_UP_RESOURCES CauseAnyOf = "INSUFFICIENT_UP_RESOURCES"
	PDU_SESSION_HANDED_OVER CauseAnyOf = "PDU_SESSION_HANDED_OVER"
	PDU_SESSION_RESUMED CauseAnyOf = "PDU_SESSION_RESUMED"
	CN_ASSISTED_RAN_PARAMETER_TUNING CauseAnyOf = "CN_ASSISTED_RAN_PARAMETER_TUNING"
	ISMF_CONTEXT_TRANSFER CauseAnyOf = "ISMF_CONTEXT_TRANSFER"
	SMF_CONTEXT_TRANSFER CauseAnyOf = "SMF_CONTEXT_TRANSFER"
	REL_DUE_TO_PS_TO_CS_HO CauseAnyOf = "REL_DUE_TO_PS_TO_CS_HO"
	REL_DUE_TO_SUBSCRIPTION_CHANGE CauseAnyOf = "REL_DUE_TO_SUBSCRIPTION_CHANGE"
	HO_CANCEL CauseAnyOf = "HO_CANCEL"
	REL_DUE_TO_SLICE_NOT_AUTHORIZED CauseAnyOf = "REL_DUE_TO_SLICE_NOT_AUTHORIZED"
	PDU_SESSION_HAND_OVER_FAILURE CauseAnyOf = "PDU_SESSION_HAND_OVER_FAILURE"
	DDN_FAILURE_STATUS CauseAnyOf = "DDN_FAILURE_STATUS"
	REL_DUE_TO_CP_ONLY_NOT_APPLICABLE CauseAnyOf = "REL_DUE_TO_CP_ONLY_NOT_APPLICABLE"
	NOT_SUPPORTED_WITH_ISMF CauseAnyOf = "NOT_SUPPORTED_WITH_ISMF"
	CHANGED_ANCHOR_SMF CauseAnyOf = "CHANGED_ANCHOR_SMF"
	CHANGED_INTERMEDIATE_SMF CauseAnyOf = "CHANGED_INTERMEDIATE_SMF"
	TARGET_DNAI_NOTIFICATION CauseAnyOf = "TARGET_DNAI_NOTIFICATION"
	REL_DUE_TO_VPLMN_QOS_FAILURE CauseAnyOf = "REL_DUE_TO_VPLMN_QOS_FAILURE"
)
