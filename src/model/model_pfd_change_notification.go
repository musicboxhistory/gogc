/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PfdChangeNotification - Represents information related to a notification of PFD change.
type PfdChangeNotification struct {

	// String providing an application identifier.
	ApplicationId string `json:"applicationId"`

	RemovalFlag bool `json:"removalFlag,omitempty"`

	PartialFlag bool `json:"partialFlag,omitempty"`

	Pfds []PfdContent `json:"pfds,omitempty"`
}
