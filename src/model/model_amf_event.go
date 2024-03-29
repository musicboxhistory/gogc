/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

import (
	"time"
)

// AmfEvent - Describes an event to be subscribed
type AmfEvent struct {
	Type AmfEventType `json:"type"`

	ImmediateFlag bool `json:"immediateFlag,omitempty"`

	AreaList []AmfEventArea `json:"areaList,omitempty"`

	LocationFilterList []LocationFilter `json:"locationFilterList,omitempty"`

	RefId int32 `json:"refId,omitempty"`

	TrafficDescriptorList []TrafficDescriptor `json:"trafficDescriptorList,omitempty"`

	ReportUeReachable bool `json:"reportUeReachable,omitempty"`

	ReachabilityFilter ReachabilityFilter `json:"reachabilityFilter,omitempty"`

	UdmDetectInd bool `json:"udmDetectInd,omitempty"`

	MaxReports int32 `json:"maxReports,omitempty"`

	// A map(list of key-value pairs) where praId serves as key.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList,omitempty"`

	// indicating a time in seconds.
	MaxResponseTime int32 `json:"maxResponseTime,omitempty"`

	TargetArea TargetArea `json:"targetArea,omitempty"`

	SnssaiFilter []ExtSnssai `json:"snssaiFilter,omitempty"`

	UeInAreaFilter UeInAreaFilter `json:"ueInAreaFilter,omitempty"`

	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	NextReport time.Time `json:"nextReport,omitempty"`

	IdleStatusInd bool `json:"idleStatusInd,omitempty"`

	DispersionArea DispersionArea `json:"dispersionArea,omitempty"`
}
