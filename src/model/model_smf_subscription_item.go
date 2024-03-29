/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmfSubscriptionItem - Contains info about a single SMF event subscription
type SmfSubscriptionItem struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfInstanceId string `json:"smfInstanceId"`

	// String providing an URI formatted according to RFC 3986
	SubscriptionId string `json:"subscriptionId"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`
}
