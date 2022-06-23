/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type LocationAccuracyAnyOf string

// List of LocationAccuracyAnyOf
const (
	CELL_LEVEL LocationAccuracyAnyOf = "CELL_LEVEL"
	RAN_NODE_LEVEL LocationAccuracyAnyOf = "RAN_NODE_LEVEL"
	TA_LEVEL LocationAccuracyAnyOf = "TA_LEVEL"
	N3_IWF_LEVEL LocationAccuracyAnyOf = "N3IWF_LEVEL"
	UE_IP LocationAccuracyAnyOf = "UE_IP"
	UE_PORT LocationAccuracyAnyOf = "UE_PORT"
)