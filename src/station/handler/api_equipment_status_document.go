/*
 * 5G-EIR Equipment Identity Check
 *
 * 5G-EIR Equipment Identity Check Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.2.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handler

import (
	"gogc/src/common/logger"
	"gogc/src/station/common"
	"gogc/src/station/equipmentstatus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetEquipmentStatus - Retrieves the status of the UE
func GetEquipmentStatus(c *gin.Context) {

	logger.Snap("GetEquipmentStatus START")
	defer logger.Snap("GetEquipmentStatus END")

	// Variable Declaration

	// Get Path Parameter
	nfType := c.Param("nfType")

	// Check NF Type
	nfType = strings.ToLower(nfType)
	ok := common.CheckNftype(nfType)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// Call Scenario Function
	output, err := equipmentstatus.Get(nfType)

	if err == nil {
		c.JSON(http.StatusOK, output)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

// PostEquipmentStatus - Retrieves the status of the UE
func PostEquipmentStatus(c *gin.Context) {

	logger.Snap("PostEquipmentStatus START")
	defer logger.Snap("PostEquipmentStatus END")
}

// PutEquipmentStatus - Retrieves the status of the UE
func PutEquipmentStatus(c *gin.Context) {

	logger.Snap("PutEquipmentStatus START")
	defer logger.Snap("PutEquipmentStatus END")
}

// DeleteEquipmentStatus - Retrieves the status of the UE
func DeleteEquipmentStatus(c *gin.Context) {

	logger.Snap("DeleteEquipmentStatus START")
	defer logger.Snap("DeleteEquipmentStatus END")
}