/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NSSAIAvailabilityDelete - Deletes an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
func NSSAIAvailabilityDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// NSSAIAvailabilityPatch - Updates an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
func NSSAIAvailabilityPatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// NSSAIAvailabilityPut - Updates/replaces the NSSF with the S-NSSAIs the NF service consumer (e.g AMF)supports per TA
func NSSAIAvailabilityPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
