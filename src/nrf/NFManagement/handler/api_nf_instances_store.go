/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNFInstances - Retrieves a collection of NF Instances
func GetNFInstances(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// OptionsNFInstances - Discover communication options supported by NRF for NF Instances
func OptionsNFInstances(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
