/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"os"
)

type ReleasePduSession200Response struct {

	JsonData ReleasedData `json:"jsonData,omitempty"`

	BinaryDataN4Information *os.File `json:"binaryDataN4Information,omitempty"`

	BinaryDataN4InformationExt1 *os.File `json:"binaryDataN4InformationExt1,omitempty"`

	BinaryDataN4InformationExt2 *os.File `json:"binaryDataN4InformationExt2,omitempty"`
}
