/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type SubscrCond struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId string `json:"nfInstanceId"`

	NfInstanceIdList []string `json:"nfInstanceIdList"`

	NfType string `json:"nfType"`

	ServiceName ServiceName `json:"serviceName"`

	ConditionType string `json:"conditionType"`

	ServiceNameList []ServiceName `json:"serviceNameList"`

	GuamiList []Guami `json:"guamiList"`

	SnssaiList []Snssai `json:"snssaiList"`

	NsiList []string `json:"nsiList,omitempty"`

	// Identifier of a group of NFs.
	NfGroupId string `json:"nfGroupId"`

	NfGroupIdList []string `json:"nfGroupIdList"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"  set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or \"set<SetID>.  <NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with <MCC> encoded as defined in clause 5.4.2  (\"Mcc\" data type definition) <MNC> encoded as defined in clause 5.4.2 (\"Mnc\" data type  definition) <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but  with lower case characters <Set ID> encoded as a string of characters consisting of alphabetic  characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an  alphabetic character or a digit.
	NfSetId string `json:"nfSetId"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string  \" set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\">\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with <MCC>  encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoded as defined in  clause 5.4.2 (\"Mnc\" data type definition)  <NID> encoded as defined in clause 5.4.2 (\"Nid\"  data type definition) <NFInstanceId> encoded as defined in clause 5.3.2 <ServiceName> encoded  as defined in 3GPP TS 29.510 <Set ID> encoded as a string of characters consisting of  alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end  with either an alphabetic character or a digit.
	NfServiceSetId string `json:"nfServiceSetId"`

	SmfServingArea []string `json:"smfServingArea,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	ScpDomains []string `json:"scpDomains"`

	NfTypeList []NfType `json:"nfTypeList,omitempty"`

	AnalyticsIds []string `json:"analyticsIds,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	ServingNfTypeList []NfType `json:"servingNfTypeList,omitempty"`

	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`

	MlAnalyticsList []MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`

	AfEvents []AfEvent `json:"afEvents,omitempty"`

	PfdData PfdData `json:"pfdData,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	ServedFqdnList []string `json:"servedFqdnList,omitempty"`
}
