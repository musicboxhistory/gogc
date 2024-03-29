/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"os"
)

type ContextUpdateRequest struct {

	JsonData ContextUpdateReqData `json:"jsonData,omitempty"`

	BinaryDataN2Information *os.File `json:"binaryDataN2Information,omitempty"`
}
