/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PccRule - Contains a PCC rule information.
type PccRule struct {

	// An array of IP flow packet filter information.
	FlowInfos []FlowInformation `json:"flowInfos,omitempty"`

	// A reference to the application detection filter configured at the UPF.
	AppId string `json:"appId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	AppDescriptor string `json:"appDescriptor,omitempty"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`

	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId string `json:"pccRuleId"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence int32 `json:"precedence,omitempty"`

	AfSigProtocol AfSigProtocol `json:"afSigProtocol,omitempty"`

	// Indication of application relocation possibility.
	AppReloc bool `json:"appReloc,omitempty"`

	// Indicates the EAS rediscovery is required.
	EasRedisInd bool `json:"easRedisInd,omitempty"`

	// A reference to the QosData policy decision type. It is the qosId described in clause 5.6.2.8.
	RefQosData []string `json:"refQosData,omitempty"`

	// A Reference to the QosData policy decision type for the Alternative QoS parameter sets of the service data flow.
	RefAltQosParams []string `json:"refAltQosParams,omitempty"`

	// A reference to the TrafficControlData policy decision type. It is the tcId described in clause 5.6.2.10.
	RefTcData []string `json:"refTcData,omitempty"`

	// A reference to the ChargingData policy decision type. It is the chgId described in clause 5.6.2.11.
	RefChgData *[]string `json:"refChgData,omitempty"`

	// A reference to the ChargingData policy decision type only applicable to Non-3GPP access if \"ATSSS\" feature is supported. It is the chgId described in clause 5.6.2.11. 
	RefChgN3gData *[]string `json:"refChgN3gData,omitempty"`

	// A reference to UsageMonitoringData policy decision type. It is the umId described in clause 5.6.2.12.
	RefUmData *[]string `json:"refUmData,omitempty"`

	// A reference to UsageMonitoringData policy decision type only applicable to Non-3GPP access if \"ATSSS\" feature is supported. It is the umId described in clause 5.6.2.12.  
	RefUmN3gData *[]string `json:"refUmN3gData,omitempty"`

	// A reference to the condition data. It is the condId described in clause 5.6.2.9.
	RefCondData *string `json:"refCondData,omitempty"`

	// A reference to the QosMonitoringData policy decision type. It is the qmId described in clause 5.6.2.40.
	RefQosMon *[]string `json:"refQosMon,omitempty"`

	AddrPreserInd *bool `json:"addrPreserInd,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom int32 `json:"tscaiTimeDom,omitempty"`

	DdNotifCtrl DownlinkDataNotificationControl `json:"ddNotifCtrl,omitempty"`

	DdNotifCtrl2 *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`

	DisUeNotif *bool `json:"disUeNotif,omitempty"`
}