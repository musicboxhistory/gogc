/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type PduSessionContinuityIndAnyOf string

// List of PduSessionContinuityIndAnyOf
const (
	MAINTAIN_PDUSESSION  PduSessionContinuityIndAnyOf = "MAINTAIN_PDUSESSION"
	RECONNECT_PDUSESSION PduSessionContinuityIndAnyOf = "RECONNECT_PDUSESSION"
	RELEASE_PDUSESSION   PduSessionContinuityIndAnyOf = "RELEASE_PDUSESSION"
)
