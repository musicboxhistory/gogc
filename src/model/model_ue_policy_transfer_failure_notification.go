/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UePolicyTransferFailureNotification - Represents information on the failure of a UE policy transfer to the UE because the UE is not reachable.
type UePolicyTransferFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause"`

	Ptis []int32 `json:"ptis"`
}
