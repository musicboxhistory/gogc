/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.2.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSubscription - Namf_EventExposure Subscribe service Operation
func CreateSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
