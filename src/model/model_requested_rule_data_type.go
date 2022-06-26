/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RequestedRuleDataType - Possible values are - CH_ID: Indicates that the requested rule data is the charging identifier.  - MS_TIME_ZONE: Indicates that the requested access network info type is the UE's timezone. - USER_LOC_INFO: Indicates that the requested access network info type is the UE's location. - RES_RELEASE: Indicates that the requested rule data is the result of the release of resource. - SUCC_RES_ALLO: Indicates that the requested rule data is the successful resource allocation. - EPS_FALLBACK: Indicates that the requested rule data is the report of QoS flow rejection due to EPS fallback.
type RequestedRuleDataType struct {
}
