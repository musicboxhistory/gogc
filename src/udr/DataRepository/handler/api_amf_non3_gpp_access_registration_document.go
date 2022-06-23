/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AmfContextNon3gpp - To modify the AMF context data of a UE using non 3gpp access in the UDR
func AmfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateAmfContextNon3gpp - To store the AMF context data of a UE using non-3gpp access in the UDR
func CreateAmfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QueryAmfContextNon3gpp - Retrieves the AMF context data of a UE using non-3gpp access
func QueryAmfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}