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

// GetOneEquipmentStatus - Retrieves the status of the UE
func GetOneEquipmentStatus(c *gin.Context) {

	logger.Debug("GetOneEquipmentStatus START")
	defer logger.Debug("GetOneEquipmentStatus END")

	// Get Path Parameter
	nfType := c.Param("nfType")
	key := c.Param("key")

	// Check NF Type
	nfType = strings.ToLower(nfType)
	ok := common.CheckNftype(nfType)
	if !ok {
		logger.Error("NF Type Error")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := equipmentstatus.GetOne(nfType, key)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

// GetEquipmentStatus - Retrieves the status of the UE
func GetEquipmentStatus(c *gin.Context) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Get Path Parameter
	nfType := c.Param("nfType")

	// Check NF Type
	nfType = strings.ToLower(nfType)
	ok := common.CheckNftype(nfType)
	if !ok {
		logger.Error("NF Type Error")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := equipmentstatus.Get(nfType)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

// PostEquipmentStatus - Retrieves the status of the UE
func PostEquipmentStatus(c *gin.Context) {

	logger.Debug("PostEquipmentStatus START")
	defer logger.Debug("PostEquipmentStatus END")

	// Variable Declaration
	var request []equipmentstatus.EquipmentStatus

	// Get Path Parameter
	nfType := c.Param("nfType")

	// Check NF Type
	nfType = strings.ToLower(nfType)
	logger.Debug("nfType:%v", nfType)
	ok := common.CheckNftype(nfType)
	if !ok {
		logger.Error("NF Type Error")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get Json Request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	logger.Debug("request:%#+v", request)

	// Call Scenario Function
	err = equipmentstatus.Post(nfType, request)

	if err == nil {
		c.JSON(http.StatusCreated, request)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
}

// PutEquipmentStatus - Retrieves the status of the UE
func PutEquipmentStatus(c *gin.Context) {

	logger.Debug("PutEquipmentStatus START")
	defer logger.Debug("PutEquipmentStatus END")
}

// DeleteEquipmentStatus - Retrieves the status of the UE
func DeleteEquipmentStatus(c *gin.Context) {

	logger.Debug("DeleteEquipmentStatus START")
	defer logger.Debug("DeleteEquipmentStatus END")

	// Get Path Parameter
	nfType := c.Param("nfType")
	key := c.Param("key")

	// Check NF Type
	nfType = strings.ToLower(nfType)
	ok := common.CheckNftype(nfType)
	if !ok {
		logger.Error("NF Type Error")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	err := equipmentstatus.Delete(nfType, key)

	if err == nil {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}
