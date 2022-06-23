/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// RefToBinaryData - This parameter provides information about the referenced binary body data.
type RefToBinaryData struct {

	// This IE shall contain the value of the Content-ID header of the referenced binary body part.
	ContentId string `json:"contentId"`
}
