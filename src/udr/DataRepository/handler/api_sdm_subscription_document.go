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

// ModifysdmSubscription - Modify an individual sdm subscription
func ModifysdmSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QuerysdmSubscription - Retrieves a individual sdmSubscription identified by subsId
func QuerysdmSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemovesdmSubscriptions - Deletes a sdmsubscriptions
func RemovesdmSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Updatesdmsubscriptions - Update an individual sdm subscriptions of a UE
func Updatesdmsubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
