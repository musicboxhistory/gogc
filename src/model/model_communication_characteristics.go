/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type CommunicationCharacteristics struct {
	PpSubsRegTimer *PpSubsRegTimer `json:"ppSubsRegTimer,omitempty"`

	PpActiveTime *PpActiveTime `json:"ppActiveTime,omitempty"`

	PpDlPacketCount *int32 `json:"ppDlPacketCount,omitempty"`

	PpDlPacketCountExt *PpDlPacketCountExt `json:"ppDlPacketCountExt,omitempty"`

	PpMaximumResponseTime *PpMaximumResponseTime `json:"ppMaximumResponseTime,omitempty"`

	PpMaximumLatency *PpMaximumLatency `json:"ppMaximumLatency,omitempty"`
}
