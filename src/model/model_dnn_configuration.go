/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type DnnConfiguration struct {
	PduSessionTypes PduSessionTypes `json:"pduSessionTypes"`

	SscModes SscModes `json:"sscModes"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	Var5gQosProfile SubscribedDefaultQos `json:"5gQosProfile,omitempty"`

	SessionAmbr Ambr `json:"sessionAmbr,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`

	StaticIpAddress []IpAddress `json:"staticIpAddress,omitempty"`

	UpSecurity UpSecurity `json:"upSecurity,omitempty"`

	PduSessionContinuityInd PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`

	// Identity of the NEF
	NiddNefId string `json:"niddNefId,omitempty"`

	NiddInfo NiddInformation `json:"niddInfo,omitempty"`

	RedundantSessionAllowed bool `json:"redundantSessionAllowed,omitempty"`

	AcsInfo AcsInfo `json:"acsInfo,omitempty"`

	Ipv4FrameRouteList []FrameRouteInfo `json:"ipv4FrameRouteList,omitempty"`

	Ipv6FrameRouteList []FrameRouteInfo `json:"ipv6FrameRouteList,omitempty"`

	AtsssAllowed bool `json:"atsssAllowed,omitempty"`

	SecondaryAuth bool `json:"secondaryAuth,omitempty"`

	DnAaaIpAddressAllocation bool `json:"dnAaaIpAddressAllocation,omitempty"`

	DnAaaAddress IpAddress `json:"dnAaaAddress,omitempty"`

	AdditionalDnAaaAddresses []IpAddress `json:"additionalDnAaaAddresses,omitempty"`

	// Fully Qualified Domain Name
	DnAaaFqdn string `json:"dnAaaFqdn,omitempty"`

	IptvAccCtrlInfo string `json:"iptvAccCtrlInfo,omitempty"`

	Ipv4Index IpIndex `json:"ipv4Index,omitempty"`

	Ipv6Index IpIndex `json:"ipv6Index,omitempty"`

	EcsAddrConfigInfo *EcsAddrConfigInfo `json:"ecsAddrConfigInfo,omitempty"`

	SharedEcsAddrConfigInfo string `json:"sharedEcsAddrConfigInfo,omitempty"`

	EasDiscoveryAuthorized bool `json:"easDiscoveryAuthorized,omitempty"`

	OnboardingInd bool `json:"onboardingInd,omitempty"`

	AerialUeInd AerialUeIndication `json:"aerialUeInd,omitempty"`
}
