/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// TrustAfInfo - Information of a trusted AF Instance
type TrustAfInfo struct {

	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`

	AfEvents []AfEvent `json:"afEvents,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	InternalGroupId []string `json:"internalGroupId,omitempty"`

	MappingInd bool `json:"mappingInd,omitempty"`
}