/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type AmfEventTypeAnyOf string

// List of AmfEventTypeAnyOf
const (
	LOCATION_REPORT                       AmfEventTypeAnyOf = "LOCATION_REPORT"
	PRESENCE_IN_AOI_REPORT                AmfEventTypeAnyOf = "PRESENCE_IN_AOI_REPORT"
	TIMEZONE_REPORT                       AmfEventTypeAnyOf = "TIMEZONE_REPORT"
	ACCESS_TYPE_REPORT                    AmfEventTypeAnyOf = "ACCESS_TYPE_REPORT"
	REGISTRATION_STATE_REPORT             AmfEventTypeAnyOf = "REGISTRATION_STATE_REPORT"
	CONNECTIVITY_STATE_REPORT             AmfEventTypeAnyOf = "CONNECTIVITY_STATE_REPORT"
	REACHABILITY_REPORT                   AmfEventTypeAnyOf = "REACHABILITY_REPORT"
	COMMUNICATION_FAILURE_REPORT          AmfEventTypeAnyOf = "COMMUNICATION_FAILURE_REPORT"
	UES_IN_AREA_REPORT                    AmfEventTypeAnyOf = "UES_IN_AREA_REPORT"
	SUBSCRIPTION_ID_CHANGE                AmfEventTypeAnyOf = "SUBSCRIPTION_ID_CHANGE"
	SUBSCRIPTION_ID_ADDITION              AmfEventTypeAnyOf = "SUBSCRIPTION_ID_ADDITION"
	LOSS_OF_CONNECTIVITY                  AmfEventTypeAnyOf = "LOSS_OF_CONNECTIVITY"
	_5_GS_USER_STATE_REPORT               AmfEventTypeAnyOf = "5GS_USER_STATE_REPORT"
	AVAILABILITY_AFTER_DDN_FAILURE        AmfEventTypeAnyOf = "AVAILABILITY_AFTER_DDN_FAILURE"
	TYPE_ALLOCATION_CODE_REPORT           AmfEventTypeAnyOf = "TYPE_ALLOCATION_CODE_REPORT"
	FREQUENT_MOBILITY_REGISTRATION_REPORT AmfEventTypeAnyOf = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	SNSSAI_TA_MAPPING_REPORT              AmfEventTypeAnyOf = "SNSSAI_TA_MAPPING_REPORT"
	UE_LOCATION_TRENDS                    AmfEventTypeAnyOf = "UE_LOCATION_TRENDS"
	UE_ACCESS_BEHAVIOR_TRENDS             AmfEventTypeAnyOf = "UE_ACCESS_BEHAVIOR_TRENDS"
	UE_MM_TRANSACTION_REPORT              AmfEventTypeAnyOf = "UE_MM_TRANSACTION_REPORT"
)
