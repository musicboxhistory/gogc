/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// SmfSubscriptionInfo - Information related to active subscriptions at the SMF(s)
type SmfSubscriptionInfo struct {
	SmfSubscriptionList []SmfSubscriptionItem `json:"smfSubscriptionList"`
}
