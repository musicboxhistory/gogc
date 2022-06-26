/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RouteSelectionParameterSet - Contains parameters that can be used to guide the Route Selection Descriptors of the URSP.
type RouteSelectionParameterSet struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence int32 `json:"precedence,omitempty"`

	// Indicates where the route selection parameters apply. It may correspond to a geographical area, for example using a geographic shape that is known to the AF and is configured by the operator to correspond to a list of or TAIs.
	SpatialValidityAreas []GeographicalArea `json:"spatialValidityAreas,omitempty"`

	// Indicates the TAIs in which the route selection parameters apply. This attribute is applicable only within the 5GC and it shall not be included in the request messages of untrusted AFs for URSP guidance.
	SpatialValidityTais []Tai1 `json:"spatialValidityTais,omitempty"`
}
