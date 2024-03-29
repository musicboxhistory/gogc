/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// GeographicArea - Geographic area specified by different shape.
type GeographicArea struct {

	Shape SupportedGadShapes `json:"shape"`

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`

	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`

	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`

	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`
}
