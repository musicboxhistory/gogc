/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AlternativeServiceRequirementsData - Contains an alternative QoS related parameter set.
type AlternativeServiceRequirementsData struct {

	// Reference to this alternative QoS related parameter set.
	AltQosParamSetRef string `json:"altQosParamSetRef"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrUl string `json:"gbrUl,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl string `json:"gbrDl,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb int32 `json:"pdb,omitempty"`
}
