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

// CreateOrUpdateNssaiAck - To store the NSSAI update acknowledgement information of a UE
func CreateOrUpdateNssaiAck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
