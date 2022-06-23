/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// FlowInfo - Represents flow information.
type FlowInfo struct {

	// Indicates the IP flow.
	FlowId int32 `json:"flowId"`

	// Indicates the packet filters of the IP flow. Refer to clause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
}