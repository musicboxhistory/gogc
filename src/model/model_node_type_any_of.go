/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type NodeTypeAnyOf string

// List of NodeTypeAnyOf
const (
	VLR             NodeTypeAnyOf = "VLR"
	SGSN            NodeTypeAnyOf = "SGSN"
	S_CSCF          NodeTypeAnyOf = "S_CSCF"
	GAN_AAA_SERVER  NodeTypeAnyOf = "GAN_AAA_SERVER"
	WLAN_AAA_SERVER NodeTypeAnyOf = "WLAN_AAA_SERVER"
)
