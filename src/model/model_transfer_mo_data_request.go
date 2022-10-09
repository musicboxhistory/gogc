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

type TransferMoDataRequest struct {

	JsonData TransferMoDataReqData `json:"jsonData,omitempty"`

	BinaryMoData *os.File `json:"binaryMoData,omitempty"`
}
