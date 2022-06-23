/*
 * Namf_MT
 *
 * AMF Mobile Terminated Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.3
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ReachabilityNotificationData - Data within the UE Reachability Info Notify
type ReachabilityNotificationData struct {

	ReachableUeList []ReachableUeInfo `json:"reachableUeList,omitempty"`

	UnreachableUeList []string `json:"unreachableUeList,omitempty"`
}
