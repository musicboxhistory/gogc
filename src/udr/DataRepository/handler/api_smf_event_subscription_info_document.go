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

// CreateSMFSubscriptions - Create SMF Subscription Info
func CreateSMFSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetSmfGroupSubscriptions - Retrieve SMF Subscription Info for a group of UEs or any UE
func GetSmfGroupSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetSmfSubscriptionInfo - Retrieve SMF Subscription Info
func GetSmfSubscriptionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifySmfGroupSubscriptions - Modify SMF Subscription Info for a group of UEs or any UE
func ModifySmfGroupSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifySmfSubscriptionInfo - Modify SMF Subscription Info
func ModifySmfSubscriptionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveSmfGroupSubscriptions - Delete SMF Subscription Info for a group of UEs or any UE
func RemoveSmfGroupSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveSmfSubscriptionsInfo - Delete SMF Subscription Info
func RemoveSmfSubscriptionsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
