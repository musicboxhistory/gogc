/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DnnSmfInfoItem - Set of parameters supported by SMF for a given DNN
type DnnSmfInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`

	DnaiList []DnnSmfInfoItemDnaiListInner `json:"dnaiList,omitempty"`
}
