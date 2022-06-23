/*
 * LMF Broadcast
 *
 * LMF Broadcast Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CipheringKeyData - Request ciphering key data
func CipheringKeyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
