/*
 * AUSF API
 *
 * AUSF UE Authentication Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// UeAuthenticationCtx - Contains the information related to the resource generated to handle the UE authentication. It contains at least the UE id, Serving Network, the Authentication Method and related EAP information or related 5G-AKA information.
type UeAuthenticationCtx struct {

	AuthType AuthType `json:"authType"`

	Var5gAuthData UeAuthenticationCtx5gAuthData `json:"5gAuthData"`

	// A map(list of key-value pairs) where member serves as key
	Links map[string]LinksValueSchema `json:"_links"`

	ServingNetworkName string `json:"servingNetworkName,omitempty"`
}