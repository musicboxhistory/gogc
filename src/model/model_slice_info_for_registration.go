/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SliceInfoForRegistration - Contains the slice information requested during a Registration procedure
type SliceInfoForRegistration struct {
	SubscribedNssai []SubscribedSnssai `json:"subscribedNssai,omitempty"`

	AllowedNssaiCurrentAccess AllowedNssai `json:"allowedNssaiCurrentAccess,omitempty"`

	AllowedNssaiOtherAccess AllowedNssai `json:"allowedNssaiOtherAccess,omitempty"`

	SNssaiForMapping []Snssai `json:"sNssaiForMapping,omitempty"`

	RequestedNssai []Snssai `json:"requestedNssai,omitempty"`

	DefaultConfiguredSnssaiInd bool `json:"defaultConfiguredSnssaiInd,omitempty"`

	MappingOfNssai []MappingOfSnssai `json:"mappingOfNssai,omitempty"`

	RequestMapping bool `json:"requestMapping,omitempty"`

	UeSupNssrgInd bool `json:"ueSupNssrgInd,omitempty"`

	SuppressNssrgInd bool `json:"suppressNssrgInd,omitempty"`

	RejectedNssaiRa []Snssai `json:"rejectedNssaiRa,omitempty"`
}
