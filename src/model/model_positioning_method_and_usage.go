/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PositioningMethodAndUsage - Indicates the usage of a positioning method.
type PositioningMethodAndUsage struct {

	Method PositioningMethod `json:"method"`

	Mode PositioningMode `json:"mode"`

	Usage Usage `json:"usage"`

	MethodCode int32 `json:"methodCode,omitempty"`
}
