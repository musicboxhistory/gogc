/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// AerialUeSubscriptionInfo - Contains the Aerial UE Subscription Information, it at least contains the Aerial UE Indication.
type AerialUeSubscriptionInfo struct {

	AerialUeInd AerialUeIndication `json:"aerialUeInd"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-<extid>, where <extid>  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Var3gppUavId string `json:"3gppUavId,omitempty"`
}
