/*
 * Namf_MBSCommunication
 *
 * AMF Communication Service for MBS. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// N2MessageTransfer - Namf_MBSCommunication N2 Message Transfer service Operation
func N2MessageTransfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
