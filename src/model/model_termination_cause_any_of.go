/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type TerminationCauseAnyOf string

// List of TerminationCauseAnyOf
const (
	ALL_SDF_DEACTIVATION                   TerminationCauseAnyOf = "ALL_SDF_DEACTIVATION"
	PDU_SESSION_TERMINATION                TerminationCauseAnyOf = "PDU_SESSION_TERMINATION"
	INSUFFICIENT_SERVER_RESOURCES          TerminationCauseAnyOf = "INSUFFICIENT_SERVER_RESOURCES"
	INSUFFICIENT_QOS_FLOW_RESOURCES        TerminationCauseAnyOf = "INSUFFICIENT_QOS_FLOW_RESOURCES"
	SPONSORED_DATA_CONNECTIVITY_DISALLOWED TerminationCauseAnyOf = "SPONSORED_DATA_CONNECTIVITY_DISALLOWED"
)
