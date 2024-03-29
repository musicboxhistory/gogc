/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DnfUnit - During the processing of dnfUnits attribute, all the members in the array shall be interpreted as logically concatenated with logical \"OR\".
type DnfUnit struct {
	DnfUnit []Atom `json:"dnfUnit"`
}
