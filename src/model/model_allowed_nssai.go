/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AllowedNssai - Contains an array of allowed S-NSSAI that constitute the allowed NSSAI information for the authorized network slice information
type AllowedNssai struct {
	AllowedSnssaiList []AllowedSnssai `json:"allowedSnssaiList"`

	AccessType AccessType `json:"accessType"`
}
