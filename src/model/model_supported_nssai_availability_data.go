/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SupportedNssaiAvailabilityData - This contains the Nssai availability data information per TA supported by the AMF
type SupportedNssaiAvailabilityData struct {

	Tai Tai `json:"tai"`

	SupportedSnssaiList []ExtSnssai `json:"supportedSnssaiList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}
