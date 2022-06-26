/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DataFilter - Identifies a data filter.
type DataFilter struct {
	DataInd DataInd `json:"dataInd"`

	Dnns []string `json:"dnns,omitempty"`

	Snssais []Snssai `json:"snssais,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	Supis []string `json:"supis,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	UeIpv4s []string `json:"ueIpv4s,omitempty"`

	UeIpv6s []Ipv6Addr `json:"ueIpv6s,omitempty"`

	UeMacs []string `json:"ueMacs,omitempty"`

	// Indicates the request is for any UE.
	AnyUeInd bool `json:"anyUeInd,omitempty"`

	// Indicates the request is for any DNN and S-NSSAI combination present in the array.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`
}
