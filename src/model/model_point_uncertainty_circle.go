/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// PointUncertaintyCircle - Ellipsoid point with uncertainty circle.
type PointUncertaintyCircle struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}
