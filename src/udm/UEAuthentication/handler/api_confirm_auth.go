/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfirmAuth - Create a new confirmation event
func ConfirmAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
