/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLcsMoData - retrieve a UE's LCS Mobile Originated Subscription Data
func GetLcsMoData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
