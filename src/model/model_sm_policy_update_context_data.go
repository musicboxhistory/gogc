/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// SmPolicyUpdateContextData - Contains the policy control request trigger(s) that were met and the corresponding new value(s) or the error report of the policy enforcement.
type SmPolicyUpdateContextData struct {

	// The policy control reqeust trigges which are met.
	RepPolicyCtrlReqTriggers []PolicyControlRequestTrigger `json:"repPolicyCtrlReqTriggers,omitempty"`

	// Indicates the access network charging identifier for the PCC rule(s) or whole PDU session.
	AccNetChIds []AccNetChId `json:"accNetChIds,omitempty"`

	AccessType AccessType `json:"accessType,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	AddAccessInfo AdditionalAccessInfo `json:"addAccessInfo,omitempty"`

	RelAccessInfo AdditionalAccessInfo `json:"relAccessInfo,omitempty"`

	ServingNetwork PlmnIdNid `json:"servingNetwork,omitempty"`

	UserLocationInfo UserLocation `json:"userLocationInfo,omitempty"`

	// String with format \"<time-numoffset>\" optionally appended by \"<daylightSavingTime>\", where  -  <time-numoffset> shall represent the time zone adjusted for daylight saving time and be encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - <daylightSavingTime> shall represent the adjustment that has been made and shall be encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.  The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone string `json:"ueTimeZone,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	RelIpv4Address string `json:"relIpv4Address,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Address string `json:"ipv4Address,omitempty"`

	// Indicates the IPv4 address domain
	IpDomain string `json:"ipDomain,omitempty"`

	Ipv6AddressPrefix Ipv6Prefix `json:"ipv6AddressPrefix,omitempty"`

	RelIpv6AddressPrefix Ipv6Prefix `json:"relIpv6AddressPrefix,omitempty"`

	AddIpv6AddrPrefixes Ipv6Prefix `json:"addIpv6AddrPrefixes,omitempty"`

	AddRelIpv6AddrPrefixes Ipv6Prefix `json:"addRelIpv6AddrPrefixes,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	RelUeMac string `json:"relUeMac,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	UeMac string `json:"ueMac,omitempty"`

	SubsSessAmbr Ambr `json:"subsSessAmbr,omitempty"`

	// Indicates the DN-AAA authorization profile index
	AuthProfIndex string `json:"authProfIndex,omitempty"`

	SubsDefQos SubscribedDefaultQos `json:"subsDefQos,omitempty"`

	VplmnQos VplmnQos `json:"vplmnQos,omitempty"`

	// If it is included and set to true, indicates that the QoS constraints in the VPLMN are not applicable.
	VplmnQosNotApp bool `json:"vplmnQosNotApp,omitempty"`

	// Contains the number of supported packet filter for signalled QoS rules.
	NumOfPackFilter int32 `json:"numOfPackFilter,omitempty"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`

	// If it is included and set to true, the 3GPP PS Data Off is activated by the UE.
	Var3gppPsDataOffStatus bool `json:"3gppPsDataOffStatus,omitempty"`

	// Report the start/stop of the application traffic and detected SDF descriptions if applicable.
	AppDetectionInfos []AppDetectionInfo `json:"appDetectionInfos,omitempty"`

	// Used to report the PCC rule failure.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`

	// Used to report the session rule failure.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`

	// QoS Notification Control information.
	QncReports []QosNotificationControlInfo `json:"qncReports,omitempty"`

	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UserLocationInfoTime time.Time `json:"userLocationInfoTime,omitempty"`

	// Reports the changes of presence reporting area. The praId attribute within the PresenceInfo data type is the key of the map. 
	RepPraInfos map[string]PresenceInfo `json:"repPraInfos,omitempty"`

	UeInitResReq UeInitiatedResourceRequest `json:"ueInitResReq,omitempty"`

	// If it is included and set to true, the reflective QoS is supported by the UE. If it is included and set to false, the reflective QoS is revoked by the UE. 
	RefQosIndication bool `json:"refQosIndication,omitempty"`

	QosFlowUsage QosFlowUsage `json:"qosFlowUsage,omitempty"`

	CreditManageStatus CreditManagementStatus `json:"creditManageStatus,omitempty"`

	ServNfId ServingNfIdentity `json:"servNfId,omitempty"`

	TraceReq *TraceData `json:"traceReq,omitempty"`

	MaPduInd MaPduIndication `json:"maPduInd,omitempty"`

	AtsssCapab AtsssCapability `json:"atsssCapab,omitempty"`

	TsnBridgeInfo TsnBridgeInfo `json:"tsnBridgeInfo,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`

	MulAddrInfos []IpMulticastAddressInfo `json:"mulAddrInfos,omitempty"`

	// Contains the type(s) of failed policy decision and/or condition data.
	PolicyDecFailureReports []PolicyDecisionFailureCode `json:"policyDecFailureReports,omitempty"`

	// Indicates the invalid parameters for the reported type(s) of the failed policy decision and/or condition data.
	InvalidPolicyDecs []InvalidParam `json:"invalidPolicyDecs,omitempty"`

	TrafficDescriptors []DddTrafficDescriptor `json:"trafficDescriptors,omitempty"`

	// Contains the identifier of the PCC rule which is used for traffic detection of event.
	PccRuleId string `json:"pccRuleId,omitempty"`

	TypesOfNotif []DlDataDeliveryStatus `json:"typesOfNotif,omitempty"`

	InterGrpIds []string `json:"interGrpIds,omitempty"`

	SatBackhaulCategory SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`

	PcfUeInfo *PcfUeCallbackInfo `json:"pcfUeInfo,omitempty"`

	NwdafDatas *[]NwdafData `json:"nwdafDatas,omitempty"`

	// When it is included and set to true, it indicates that the AN-Gateway has failed and that the PCF should refrain from sending policy decisions to the SMF until it is informed that the AN-Gateway has been recovered. 
	AnGwStatus bool `json:"anGwStatus,omitempty"`
}