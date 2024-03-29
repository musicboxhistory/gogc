/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ProSeAllowedPlmn - Contains the PLMN identities where the Prose services are authorised to use and the authorised Prose services on this given PLMNs.
type ProSeAllowedPlmn struct {
	VisitedPlmn PlmnId `json:"visitedPlmn"`

	ProseDirectAllowed []ProseDirectAllowed `json:"proseDirectAllowed,omitempty"`
}
