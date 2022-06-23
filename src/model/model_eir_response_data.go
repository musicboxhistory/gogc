/*
 * 5G-EIR Equipment Identity Check
 *
 * 5G-EIR Equipment Identity Check Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// EirResponseData - Represents equipment status data provided in an EIR response message.
type EirResponseData struct {

	Status EquipmentStatus `json:"status"`
}