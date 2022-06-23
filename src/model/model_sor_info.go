/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

type SorInfo struct {

	SteeringContainer SteeringContainer `json:"steeringContainer,omitempty"`

	// Contains indication whether the acknowledgement from UE is needed.
	AckInd bool `json:"ackInd"`

	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIausf string `json:"sorMacIausf,omitempty"`

	// CounterSoR.
	Countersor string `json:"countersor,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	// string with format 'bytes' as defined in OpenAPI
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	SorCmci string `json:"sorCmci,omitempty"`

	StoreSorCmciInMe bool `json:"storeSorCmciInMe,omitempty"`

	UsimSupportOfSorCmci bool `json:"usimSupportOfSorCmci,omitempty"`
}
