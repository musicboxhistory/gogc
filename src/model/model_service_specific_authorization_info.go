/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// ServiceSpecificAuthorizationInfo - Information related to active Service Specific Authorizations
type ServiceSpecificAuthorizationInfo struct {

	ServiceSpecificAuthorizationList []AuthorizationInfo `json:"serviceSpecificAuthorizationList"`
}