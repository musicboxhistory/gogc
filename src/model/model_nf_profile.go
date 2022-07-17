/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// NfProfile - Information of an NF Instance registered in the NRF
type NfProfile struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId string `json:"nfInstanceId" bson:"nfInstanceId"`

	NfInstanceName *string `json:"nfInstanceName,omitempty" bson:"nfInstanceName,omitempty"`

	NfType *NfType `json:"nfType" bson:"nfType"`

	NfStatus *NfStatus `json:"nfStatus" bson:"nfStatus"`

	CollocatedNfInstances *[]CollocatedNfInstance `json:"collocatedNfInstances,omitempty" bson:"collocatedNfInstances,omitempty"`

	HeartBeatTimer *int32 `json:"heartBeatTimer,omitempty" bson:"heartBeatTimer,omitempty"`

	PlmnList *[]PlmnId `json:"plmnList,omitempty" bson:"plmnList,omitempty"`

	SnpnList *[]PlmnIdNid `json:"snpnList,omitempty" bson:"snpnList,omitempty"`

	SNssais *[]ExtSnssai `json:"sNssais,omitempty" bson:"sNssais,omitempty"`

	PerPlmnSnssaiList *[]PlmnSnssai `json:"perPlmnSnssaiList,omitempty" bson:"perPlmnSnssaiList,omitempty"`

	NsiList *[]string `json:"nsiList,omitempty" bson:"nsiList,omitempty"`

	// Fully Qualified Domain Name
	Fqdn *string `json:"fqdn,omitempty" bson:"fqdn,omitempty"`

	// Fully Qualified Domain Name
	InterPlmnFqdn *string `json:"interPlmnFqdn,omitempty" bson:"interPlmnFqdn,omitempty"`

	Ipv4Addresses *[]string `json:"ipv4Addresses,omitempty" bson:"ipv4Addresses,omitempty"`

	Ipv6Addresses *[]Ipv6Addr `json:"ipv6Addresses,omitempty" bson:"ipv6Addresses,omitempty"`

	AllowedPlmns *[]PlmnId `json:"allowedPlmns,omitempty" bson:"allowedPlmns,omitempty"`

	AllowedSnpns *[]PlmnIdNid `json:"allowedSnpns,omitempty" bson:"allowedSnpns,omitempty"`

	AllowedNfTypes *[]NfType `json:"allowedNfTypes,omitempty" bson:"allowedNfTypes,omitempty"`

	AllowedNfDomains *[]string `json:"allowedNfDomains,omitempty" bson:"allowedNfDomains,omitempty"`

	AllowedNssais *[]ExtSnssai `json:"allowedNssais,omitempty" bson:"allowedNssais,omitempty"`

	Priority *int32 `json:"priority,omitempty" bson:"priority,omitempty"`

	Capacity *int32 `json:"capacity,omitempty" bson:"capacity,omitempty"`

	Load *int32 `json:"load,omitempty" bson:"load,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp *time.Time `json:"loadTimeStamp,omitempty" bson:"loadTimeStamp,omitempty"`

	Locality *string `json:"locality,omitempty" bson:"locality,omitempty"`

	UdrInfo *UdrInfo `json:"udrInfo,omitempty" bson:"udrInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdrInfo
	UdrInfoList map[string]UdrInfo `json:"udrInfoList,omitempty" bson:"udrInfoList,omitempty"`

	UdmInfo *UdmInfo `json:"udmInfo,omitempty" bson:"udmInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdmInfo
	UdmInfoList map[string]UdmInfo `json:"udmInfoList,omitempty" bson:"udmInfoList,omitempty"`

	AusfInfo *AusfInfo `json:"ausfInfo,omitempty" bson:"ausfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AusfInfo
	AusfInfoList map[string]AusfInfo `json:"ausfInfoList,omitempty" bson:"ausfInfoList,omitempty"`

	AmfInfo *AmfInfo `json:"amfInfo,omitempty" bson:"amfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AmfInfo
	AmfInfoList map[string]AmfInfo `json:"amfInfoList,omitempty" bson:"amfInfoList,omitempty"`

	SmfInfo *SmfInfo `json:"smfInfo,omitempty" bson:"smfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of SmfInfo
	SmfInfoList map[string]SmfInfo `json:"smfInfoList,omitempty" bson:"smfInfoList,omitempty"`

	UpfInfo *UpfInfo `json:"upfInfo,omitempty" bson:"upfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UpfInfo
	UpfInfoList map[string]UpfInfo `json:"upfInfoList,omitempty" bson:"upfInfoList,omitempty"`

	PcfInfo *PcfInfo `json:"pcfInfo,omitempty" bson:"pcfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcfInfo
	PcfInfoList map[string]PcfInfo `json:"pcfInfoList,omitempty" bson:"pcfInfoList,omitempty"`

	BsfInfo *BsfInfo `json:"bsfInfo,omitempty" bson:"bsfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of BsfInfo
	BsfInfoList map[string]BsfInfo `json:"bsfInfoList,omitempty" bson:"bsfInfoList,omitempty"`

	ChfInfo *ChfInfo `json:"chfInfo,omitempty" bson:"chfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ChfInfo
	ChfInfoList map[string]ChfInfo `json:"chfInfoList,omitempty" bson:"chfInfoList,omitempty"`

	NefInfo *NefInfo `json:"nefInfo,omitempty" bson:"nefInfo,omitempty"`

	NrfInfo *NrfInfo `json:"nrfInfo,omitempty" bson:"nrfInfo,omitempty"`

	UdsfInfo *UdsfInfo `json:"udsfInfo,omitempty" bson:"udsfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdsfInfo
	UdsfInfoList map[string]UdsfInfo `json:"udsfInfoList,omitempty" bson:"udsfInfoList,omitempty"`

	NwdafInfo *NwdafInfo `json:"nwdafInfo,omitempty" bson:"nwdafInfo,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	NwdafInfoList map[string]NwdafInfo `json:"nwdafInfoList,omitempty" bson:"nwdafInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcscfInfo
	PcscfInfoList map[string]PcscfInfo `json:"pcscfInfoList,omitempty" bson:"pcscfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of HssInfo
	HssInfoList map[string]HssInfo `json:"hssInfoList,omitempty" bson:"hssInfoList,omitempty"`

	CustomInfo map[string]interface{} `json:"customInfo,omitempty" bson:"customInfo,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty" json:"recoveryTime,omitempty"`

	NfServicePersistence *bool `json:"nfServicePersistence,omitempty" bson:"nfServicePersistence,omitempty"`

	// Deprecated
	NfServices *[]NfService `json:"nfServices,omitempty" bson:"nfServices,omitempty"`

	// A map (list of key-value pairs) where serviceInstanceId serves as key of NFService
	NfServiceList map[string]NfService `json:"nfServiceList,omitempty" bson:"nfServiceList,omitempty"`

	NfProfileChangesSupportInd *bool `json:"nfProfileChangesSupportInd,omitempty" bson:"nfProfileChangesSupportInd,omitempty"`

	NfProfileChangesInd *bool `json:"nfProfileChangesInd,omitempty" bson:"nfProfileChangesInd,omitempty"`

	DefaultNotificationSubscriptions *[]DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty" bson:"defaultNotificationSubscriptions,omitempty"`

	LmfInfo *LmfInfo `json:"lmfInfo,omitempty" bson:"lmfInfo,omitempty"`

	GmlcInfo *GmlcInfo `json:"gmlcInfo,omitempty" bson:"gmlcInfo,omitempty"`

	NfSetIdList *[]string `json:"nfSetIdList,omitempty" bson:"nfSetIdList,omitempty"`

	ServingScope *[]string `json:"servingScope,omitempty" bson:"servingScope,omitempty"`

	LcHSupportInd *bool `json:"lcHSupportInd,omitempty" bson:"lcHSupportInd,omitempty"`

	OlcHSupportInd *bool `json:"olcHSupportInd,omitempty" bson:"olcHSupportInd,omitempty"`

	// A map (list of key-value pairs) where NfSetId serves as key of DateTime
	NfSetRecoveryTimeList map[string]time.Time `json:"nfSetRecoveryTimeList,omitempty" bson:"nfSetRecoveryTimeList,omitempty"`

	// A map (list of key-value pairs) where NfServiceSetId serves as key of DateTime
	ServiceSetRecoveryTimeList map[string]time.Time `json:"serviceSetRecoveryTimeList,omitempty" bson:"serviceSetRecoveryTimeList,omitempty"`

	ScpDomains *[]string `json:"scpDomains,omitempty" bson:"scpDomains,omitempty"`

	ScpInfo *ScpInfo `json:"scpInfo,omitempty" bson:"scpInfo,omitempty"`

	SeppInfo *SeppInfo `json:"seppInfo,omitempty" bson:"seppInfo,omitempty"`

	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId *string `json:"vendorId,omitempty" bson:"vendorId,omitempty"`

	// the key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes
	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty" bson:"supportedVendorSpecificFeatures,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	AanfInfoList map[string]AanfInfo `json:"aanfInfoList,omitempty" bson:"aanfInfoList,omitempty"`

	Var5gDdnmfInfo *Model5GDdnmfInfo `json:"5gDdnmfInfo,omitempty" bson:"5gDdnmfInfo,omitempty"`

	MfafInfo *MfafInfo `json:"mfafInfo,omitempty" bson:"mfafInfo,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	EasdfInfoList map[string]EasdfInfo `json:"easdfInfoList,omitempty" bson:"easdfInfoList,omitempty"`

	DccfInfo *DccfInfo `json:"dccfInfo,omitempty" bson:"dccfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of nsacfInfo
	NsacfInfoList map[string]NsacfInfo `json:"nsacfInfoList,omitempty" bson:"nsacfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbSmfInfo
	MbSmfInfoList map[string]MbSmfInfo `json:"mbSmfInfoList,omitempty" bson:"mbSmfInfoList,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	TsctsfInfoList map[string]TsctsfInfo `json:"tsctsfInfoList,omitempty" bson:"mbSmfInfoList,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	MbUpfInfoList map[string]MbUpfInfo `json:"mbUpfInfoList,omitempty" bson:"mbUpfInfoList,omitempty"`

	TrustAfInfo *TrustAfInfo `json:"trustAfInfo,omitempty" bson:"trustAfInfo,omitempty"`

	NssaafInfo *NssaafInfo `json:"nssaafInfo,omitempty" bson:"nssaafInfo,omitempty"`

	HniList *[]string `json:"hniList,omitempty" bson:"hniList,omitempty"`
}
