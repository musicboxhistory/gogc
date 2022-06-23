/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ComplexQuery - The ComplexQuery data type is either a conjunctive normal form or a disjunctive normal form. The attribute names \"cnfUnits\" and \"dnfUnits\" (see clause 5.2.4.11 and clause 5.2.4.12) serve as discriminator.
type ComplexQuery struct {

	CnfUnits []CnfUnit `json:"cnfUnits"`

	DnfUnits []DnfUnit `json:"dnfUnits"`
}
