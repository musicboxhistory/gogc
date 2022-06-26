/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type SessionManagementSubscriptionData1 struct {
	SingleNssai Snssai `json:"singleNssai"`

	// A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations
	DnnConfigurations map[string]DnnConfiguration1 `json:"dnnConfigurations,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	// A map(list of key-value pairs) where GroupId serves as key of SharedDataId
	SharedVnGroupDataIds map[string]string `json:"sharedVnGroupDataIds,omitempty"`

	SharedDnnConfigurationsId string `json:"sharedDnnConfigurationsId,omitempty"`

	OdbPacketServices OdbPacketServices `json:"odbPacketServices,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	SharedTraceDataId string `json:"sharedTraceDataId,omitempty"`

	// A map(list of key-value pairs) where Dnn serves as key of ExpectedUeBehaviourData
	ExpectedUeBehavioursList map[string]ExpectedUeBehaviourData1 `json:"expectedUeBehavioursList,omitempty"`

	// A map(list of key-value pairs) where Dnn serves as key of SuggestedPacketNumDl
	SuggestedPacketNumDlList map[string]SuggestedPacketNumDl1 `json:"suggestedPacketNumDlList,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`
}
