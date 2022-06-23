/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model
// SmsDeliveryStatus : Represents the Status of an SMS delivery attempt.
type SmsDeliveryStatus string

// List of SmsDeliveryStatus
const (
	PENDING SmsDeliveryStatus = "SMS_DELIVERY_PENDING"
	COMPLETED SmsDeliveryStatus = "SMS_DELIVERY_COMPLETED"
	FAILED SmsDeliveryStatus = "SMS_DELIVERY_FAILED"
	SMSF_ACCEPTED SmsDeliveryStatus = "SMS_DELIVERY_SMSF_ACCEPTED"
)
