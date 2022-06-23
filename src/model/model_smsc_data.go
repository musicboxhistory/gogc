/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmscData - Addresses of SM-Service Center entities with SMS wating to be delivered to the UE.
type SmscData struct {

	SmscMapAddress string `json:"smscMapAddress,omitempty"`

	SmscDiameterAddress NetworkNodeDiameterAddress1 `json:"smscDiameterAddress,omitempty"`
}