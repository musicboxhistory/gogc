/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://model-generator.tech)
 */

package model

type ProseDirectAllowedAnyOf string

// List of ProseDirectAllowedAnyOf
const (
	ANNOUNCE           ProseDirectAllowedAnyOf = "ANNOUNCE"
	MONITOR            ProseDirectAllowedAnyOf = "MONITOR"
	RESTRICTD_ANNOUNCE ProseDirectAllowedAnyOf = "RESTRICTD_ANNOUNCE"
	RESTRICTD_MONITOR  ProseDirectAllowedAnyOf = "RESTRICTD_MONITOR"
	DISCOVERER         ProseDirectAllowedAnyOf = "DISCOVERER"
	DISCOVEREE         ProseDirectAllowedAnyOf = "DISCOVEREE"
	BROADCAST          ProseDirectAllowedAnyOf = "BROADCAST"
	GROUPCAST          ProseDirectAllowedAnyOf = "GROUPCAST"
	UNICAST            ProseDirectAllowedAnyOf = "UNICAST"
	LAYER2_RELAY       ProseDirectAllowedAnyOf = "LAYER2_RELAY"
	LAYER3_RELAY       ProseDirectAllowedAnyOf = "LAYER3_RELAY"
)
