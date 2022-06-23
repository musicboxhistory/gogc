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

// CreateSdmSubscriptions - Create individual sdm subscription
func CreateSdmSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Querysdmsubscriptions - Retrieves the sdm subscriptions of a UE
func Querysdmsubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
