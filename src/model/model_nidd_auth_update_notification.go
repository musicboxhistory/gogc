/*
 * Nudm_NIDDAU
 *
 * Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0-alpha.3
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// NiddAuthUpdateNotification - Represents a NIDD authorization update notification.
type NiddAuthUpdateNotification struct {

	NiddAuthUpdateInfoList []NiddAuthUpdateInfo `json:"niddAuthUpdateInfoList"`
}