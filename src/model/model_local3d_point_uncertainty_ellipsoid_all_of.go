/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type Local3dPointUncertaintyEllipsoidAllOf struct {
	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
