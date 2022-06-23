/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// NsiInformation - Contains the API URIs of NRF services to be used to discover NFs/services, subscribe to NF status changes and/or request access tokens within the selected Network Slice instance and optional the Identifier of the selected Network Slice instance 
type NsiInformation struct {

	// String providing an URI formatted according to RFC 3986
	NrfId string `json:"nrfId"`

	// Contains the Identifier of the selected Network Slice instance
	NsiId string `json:"nsiId,omitempty"`

	// String providing an URI formatted according to RFC 3986
	NrfNfMgtUri string `json:"nrfNfMgtUri,omitempty"`

	// String providing an URI formatted according to RFC 3986
	NrfAccessTokenUri string `json:"nrfAccessTokenUri,omitempty"`

	// Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \"nnrf-nfm\" or \"nnrf-disc\" 
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`
}
