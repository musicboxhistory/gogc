/*
 * LMF Location
 *
 * LMF Location Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type GnssIdAnyOf string

// List of GnssIdAnyOf
const (
	GPS GnssIdAnyOf = "GPS"
	GALILEO GnssIdAnyOf = "GALILEO"
	SBAS GnssIdAnyOf = "SBAS"
	MODERNIZED_GPS GnssIdAnyOf = "MODERNIZED_GPS"
	QZSS GnssIdAnyOf = "QZSS"
	GLONASS GnssIdAnyOf = "GLONASS"
	BDS GnssIdAnyOf = "BDS"
	NAVIC GnssIdAnyOf = "NAVIC"
)
