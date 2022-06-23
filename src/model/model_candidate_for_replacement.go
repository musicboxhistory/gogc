/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// CandidateForReplacement - Represents a list of candidate DNNs for replacement for an S-NSSAI.
type CandidateForReplacement struct {

	Snssai Snssai `json:"snssai"`

	Dnns *[]string `json:"dnns,omitempty"`
}