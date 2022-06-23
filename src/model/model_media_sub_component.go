/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// MediaSubComponent - Identifies a media subcomponent.
type MediaSubComponent struct {

	AfSigProtocol AfSigProtocol `json:"afSigProtocol,omitempty"`

	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty"`

	FNum int32 `json:"fNum"`

	FDescs []string `json:"fDescs,omitempty"`

	FStatus FlowStatus `json:"fStatus,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl string `json:"marBwDl,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl string `json:"marBwUl,omitempty"`

	// 2-octet string, where each octet is encoded in hexadecimal representation. The first octet contains the IPv4 Type-of-Service or the IPv6 Traffic-Class field and the second octet contains the ToS/Traffic Class mask field. 
	TosTrCl string `json:"tosTrCl,omitempty"`

	FlowUsage FlowUsage `json:"flowUsage,omitempty"`
}
