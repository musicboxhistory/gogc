/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// SubscriptionData - Information of a subscription to notifications to NRF events, included in subscription requests and responses
type SubscriptionData struct {
	NfStatusNotificationUri string `json:"nfStatusNotificationUri" bson:"nfStatusNotificationUri"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ReqNfInstanceId *string `json:"reqNfInstanceId,omitempty" bson:"reqNfInstanceId,omitempty"`

	SubscrCond *SubscrCond `json:"subscrCond,omitempty" bson:"subscrCond,omitempty"`

	SubscriptionId string `json:"subscriptionId" bson:"subscriptionId"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty" bson:"validityTime,omitempty"`

	ReqNotifEvents *[]NotificationEventType `json:"reqNotifEvents,omitempty" bson:"reqNotifEvents,omitempty"`

	PlmnId *PlmnId `json:"plmnId,omitempty" bson:"plmnId,omitempty"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty" bson:"nid,omitempty"`

	NotifCondition *NotifCondition `json:"notifCondition,omitempty" bson:"notifCondition,omitempty"`

	ReqNfType *NfType `json:"reqNfType,omitempty" bson:"reqNfType,omitempty"`

	// Fully Qualified Domain Name
	ReqNfFqdn *string `json:"reqNfFqdn,omitempty" bson:"reqNfFqdn,omitempty"`

	ReqSnssais *[]Snssai `json:"reqSnssais,omitempty" bson:"reqSnssais,omitempty"`

	ReqPerPlmnSnssais *[]PlmnSnssai `json:"reqPerPlmnSnssais,omitempty" bson:"reqPerPlmnSnssais,omitempty"`

	ReqPlmnList *[]PlmnId `json:"reqPlmnList,omitempty" bson:"reqPlmnList,omitempty"`

	ReqSnpnList *[]PlmnIdNid `json:"reqSnpnList,omitempty" bson:"reqSnpnList,omitempty"`

	ServingScope *[]string `json:"servingScope,omitempty" bson:"servingScope,omitempty"`

	RequesterFeatures *string `json:"requesterFeatures,omitempty" bson:"requesterFeatures,omitempty"`

	NrfSupportedFeatures *string `json:"nrfSupportedFeatures,omitempty" bson:"nrfSupportedFeatures,omitempty"`

	// String providing an URI formatted according to RFC 3986
	HnrfUri *string `json:"hnrfUri,omitempty" bson:"hnrfUri,omitempty"`

	OnboardingCapability *bool `json:"onboardingCapability,omitempty" bson:"onboardingCapability,omitempty"`

	// Fully Qualified Domain Name
	TargetHni *string `json:"targetHni,omitempty" bson:"targetHni,omitempty"`
}
