/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SnssaiUpfInfoItem - Set of parameters supported by UPF for a given S-NSSAI
type SnssaiUpfInfoItem struct {
	SNssai Snssai `json:"sNssai"`

	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList"`

	RedundantTransport bool `json:"redundantTransport,omitempty"`
}
