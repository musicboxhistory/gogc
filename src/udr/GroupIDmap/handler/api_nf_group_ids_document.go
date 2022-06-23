/*
 * Nudr_GroupIDmap
 *
 * Unified Data Repository Service for NF-Group ID retrieval. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0-alpha.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNfGroupIDs - Retrieves NF-Group IDs for provided Subscriber and NF types
func GetNfGroupIDs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}