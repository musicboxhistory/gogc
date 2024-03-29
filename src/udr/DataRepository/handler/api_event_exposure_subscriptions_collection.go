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

// CreateEeSubscriptions - Create individual EE subscription
func CreateEeSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Queryeesubscriptions - Retrieves the ee subscriptions of a UE
func Queryeesubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
