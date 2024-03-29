/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PfdContent - Represents the content of a PFD for an application identifier.
type PfdContent struct {

	// Identifies a PDF of an application identifier.
	PfdId string `json:"pfdId,omitempty"`

	// Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`

	// Indicates a URL or a regular expression which is used to match the significant parts of the URL.
	Urls []string `json:"urls,omitempty"`

	// Indicates an FQDN or a regular expression as a domain name matching criteria.
	DomainNames []string `json:"domainNames,omitempty"`

	DnProtocol DomainNameProtocol `json:"dnProtocol,omitempty"`
}
