/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// MdtConfiguration - contains contain MDT configuration data.
type MdtConfiguration struct {

	JobType JobType `json:"jobType"`

	ReportType ReportTypeMdt `json:"reportType,omitempty"`

	AreaScope AreaScope `json:"areaScope,omitempty"`

	MeasurementLteList []MeasurementLteForMdt `json:"measurementLteList,omitempty"`

	MeasurementNrList []MeasurementNrForMdt `json:"measurementNrList,omitempty"`

	SensorMeasurementList []SensorMeasurement `json:"sensorMeasurementList,omitempty"`

	ReportingTriggerList []ReportingTrigger `json:"reportingTriggerList,omitempty"`

	ReportInterval ReportIntervalMdt `json:"reportInterval,omitempty"`

	ReportIntervalNr ReportIntervalNrMdt `json:"reportIntervalNr,omitempty"`

	ReportAmount ReportAmountMdt `json:"reportAmount,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-97. 
	EventThresholdRsrp int32 `json:"eventThresholdRsrp,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR.When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-127. 
	EventThresholdRsrpNr int32 `json:"eventThresholdRsrpNr,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-34. 
	EventThresholdRsrq int32 `json:"eventThresholdRsrq,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-127. 
	EventThresholdRsrqNr int32 `json:"eventThresholdRsrqNr,omitempty"`

	EventList []EventForMdt `json:"eventList,omitempty"`

	LoggingInterval LoggingIntervalMdt `json:"loggingInterval,omitempty"`

	LoggingIntervalNr LoggingIntervalNrMdt `json:"loggingIntervalNr,omitempty"`

	LoggingDuration LoggingDurationMdt `json:"loggingDuration,omitempty"`

	LoggingDurationNr LoggingDurationNrMdt `json:"loggingDurationNr,omitempty"`

	PositioningMethod PositioningMethodMdt `json:"positioningMethod,omitempty"`

	AddPositioningMethodList []PositioningMethodMdt `json:"addPositioningMethodList,omitempty"`

	CollectionPeriodRmmLte CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`

	CollectionPeriodRmmNr CollectionPeriodRmmNrMdt `json:"collectionPeriodRmmNr,omitempty"`

	MeasurementPeriodLte MeasurementPeriodLteMdt `json:"measurementPeriodLte,omitempty"`

	MdtAllowedPlmnIdList []PlmnId `json:"mdtAllowedPlmnIdList,omitempty"`

	MbsfnAreaList []MbsfnArea `json:"mbsfnAreaList,omitempty"`

	InterFreqTargetList []InterFreqTargetInfo `json:"interFreqTargetList,omitempty"`
}
