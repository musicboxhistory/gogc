/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://handler-generator.tech)
 */

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSmsfContextNon3gpp - Create the SMSF context data of a UE via non-3GPP access
func CreateSmsfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteSmsfContextNon3gpp - To remove the SMSF context data of a UE via non-3GPP access
func DeleteSmsfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QuerySmsfContextNon3gpp - Retrieves the SMSF context data of a UE using non-3gpp access
func QuerySmsfContextNon3gpp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
