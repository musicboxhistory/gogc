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

type ModifysubscriptionDataSubscription200Response struct {

	// The execution report contains an array of report items. Each report item indicates one failed modification.
	Report []ReportItem `json:"report"`

	// String represents the SUPI or GPSI
	UeId string `json:"ueId,omitempty"`

	// String providing an URI formatted according to RFC 3986
	CallbackReference string `json:"callbackReference"`

	// String providing an URI formatted according to RFC 3986
	OriginalCallbackReference string `json:"originalCallbackReference,omitempty"`

	MonitoredResourceUris []string `json:"monitoredResourceUris"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	SdmSubscription SdmSubscription1 `json:"sdmSubscription,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	UniqueSubscription bool `json:"uniqueSubscription,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
