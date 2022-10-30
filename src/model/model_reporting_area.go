/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ReportingArea - Indicates an area for event reporting.
type ReportingArea struct {

	AreaType ReportingAreaType `json:"areaType"`

	Tai Tai `json:"tai,omitempty"`

	Ecgi Ecgi `json:"ecgi,omitempty"`

	Ncgi Ncgi `json:"ncgi,omitempty"`
}