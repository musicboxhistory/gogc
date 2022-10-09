/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type SbiBindingLevelAnyOf string

// List of SbiBindingLevelAnyOf
const (
	INSTANCE_BINDING SbiBindingLevelAnyOf = "NF_INSTANCE_BINDING"
	SET_BINDING SbiBindingLevelAnyOf = "NF_SET_BINDING"
	SERVICE_SET_BINDING SbiBindingLevelAnyOf = "NF_SERVICE_SET_BINDING"
	SERVICE_INSTANCE_BINDING SbiBindingLevelAnyOf = "NF_SERVICE_INSTANCE_BINDING"
)
