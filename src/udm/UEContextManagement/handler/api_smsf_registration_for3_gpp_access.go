/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Call3GppSmsfRegistration - register as SMSF for 3GPP access
func Call3GppSmsfRegistration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
