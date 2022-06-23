/*
 * Naf_EventExposure
 *
 * AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// EventFilter - Represents event filter information for an event.
type EventFilter struct {

	Gpsis []string `json:"gpsis,omitempty"`

	Supis []string `json:"supis,omitempty"`

	ExterGroupIds []string `json:"exterGroupIds,omitempty"`

	InterGroupIds []string `json:"interGroupIds,omitempty"`

	AnyUeInd bool `json:"anyUeInd,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	LocArea LocationArea5G `json:"locArea,omitempty"`

	CollAttrs []CollectiveBehaviourFilter `json:"collAttrs,omitempty"`
}
