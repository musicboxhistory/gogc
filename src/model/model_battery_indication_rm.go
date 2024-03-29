/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// BatteryIndicationRm - This data type is defined in the same way as the 'BatteryIndication' data type, but with the OpenAPI 'nullable: true' property.
type BatteryIndicationRm struct {

	// This IE shall indicate whether the UE is battery powered or not. true: the UE is battery powered; false or absent: the UE is not battery powered
	BatteryInd bool `json:"batteryInd,omitempty"`

	// This IE shall indicate whether the battery of the UE is replaceable or not. true: the battery of the UE is replaceable; false or absent: the battery of the UE is not replaceable.
	ReplaceableInd bool `json:"replaceableInd,omitempty"`

	// This IE shall indicate whether the battery of the UE is rechargeable or not. true: the battery of UE is rechargeable; false or absent: the battery of the UE is not rechargeable.
	RechargeableInd bool `json:"rechargeableInd,omitempty"`
}
