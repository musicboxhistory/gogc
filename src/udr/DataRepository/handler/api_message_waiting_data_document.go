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

// CreateMessageWaitingData - Create the Message Waiting Data of the UE
func CreateMessageWaitingData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteMessageWaitingData - To remove the Message Waiting Data of the UE
func DeleteMessageWaitingData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifyMessageWaitingData - Modify the Message Waiting Data of the UE
func ModifyMessageWaitingData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QueryMessageWaitingData - Retrieves the Message Waiting Data of the UE
func QueryMessageWaitingData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
