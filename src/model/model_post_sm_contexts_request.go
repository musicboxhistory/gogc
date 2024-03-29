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

type PostSmContextsRequest struct {

	JsonData SmContextCreateData `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage *os.File `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmInformation *os.File `json:"binaryDataN2SmInformation,omitempty"`

	BinaryDataN2SmInformationExt1 *os.File `json:"binaryDataN2SmInformationExt1,omitempty"`
}
