/*
 * Namf_MBSCommunication
 *
 * AMF Communication Service for MBS. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// N2MbsSmInfo - N2 MBS Session Management information
type N2MbsSmInfo struct {
	NgapIeType MbsNgapIeType `json:"ngapIeType"`

	NgapData RefToBinaryData `json:"ngapData"`
}
