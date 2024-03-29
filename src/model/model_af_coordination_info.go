/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AfCoordinationInfo - AF Coordination Information
type AfCoordinationInfo struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai string `json:"sourceDnai,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SourceUeIpv4Addr string `json:"sourceUeIpv4Addr,omitempty"`

	SourceUeIpv6Prefix Ipv6Prefix `json:"sourceUeIpv6Prefix,omitempty"`

	NotificationInfoList []NotificationInfo `json:"notificationInfoList,omitempty"`
}
