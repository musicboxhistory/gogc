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

type Nssai struct {

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`

	DefaultSingleNssais *[]Snssai `json:"defaultSingleNssais"`

	SingleNssais *[]Snssai `json:"singleNssais,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime *time.Time `json:"provisioningTime,omitempty"`

	// A map(list of key-value pairs) where singleNssai serves as key of AdditionalSnssaiData
	AdditionalSnssaiData map[string]AdditionalSnssaiData `json:"additionalSnssaiData,omitempty"`

	SuppressNssrgInd *bool `json:"suppressNssrgInd,omitempty"`
}
