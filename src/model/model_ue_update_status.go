/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UeUpdateStatus : Status of the procedure.
type UeUpdateStatus string

// List of UeUpdateStatus
const (
	NOT_SENT              UeUpdateStatus = "NOT_SENT"
	SENT_NO_ACK_REQUIRED  UeUpdateStatus = "SENT_NO_ACK_REQUIRED"
	WAITING_FOR_ACK       UeUpdateStatus = "WAITING_FOR_ACK"
	ACK_RECEIVED          UeUpdateStatus = "ACK_RECEIVED"
	NEGATIVE_ACK_RECEIVED UeUpdateStatus = "NEGATIVE_ACK_RECEIVED"
)
