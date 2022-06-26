/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

// DlDataDeliveryStatus - Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is transmitted after previous buffering or discarding of corresponding packet(s) because the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is discarded because the Extended Buffering time, as determined by the SMF, expires or the amount of downlink data to be buffered is exceeded.
type DlDataDeliveryStatus struct {
}
