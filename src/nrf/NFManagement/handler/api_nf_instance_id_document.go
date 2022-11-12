/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"gogc/src/common/logger"
	"gogc/src/common/signal"
	"gogc/src/model"
	"gogc/src/nrf/NFManagement/scenario"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNFInstance - Read the profile of a given NF Instance
func GetNFInstance(c *gin.Context) {

	logger.Debug("GetNFInstance START")
	defer logger.Debug("GetNFInstance END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["nfInstanceID"] = c.Param("nfInstanceID")
	logger.Debug("request:%#+v", request)

	// Call Scenario Function
	response, err := scenario.GetNFInstance(request)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}

// RegisterNFInstance - Register a new NF Instance
func RegisterNFInstance(c *gin.Context) {

	logger.Debug("RegisterNFInstance START")
	defer logger.Debug("RegisterNFInstance END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["nfInstanceID"] = c.Param("nfInstanceID")
	logger.Debug("request:%#+v", request)

	nfProfile := model.NfProfile{}
	err := c.ShouldBindJSON(&nfProfile)
	if err != nil {
		logger.Error("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := scenario.RegisterNFInstance(request, &nfProfile)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusCreated, nfProfile)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
}

// UpdateNFInstance - Update NF Instance profile
func UpdateNFInstance(c *gin.Context) {

	logger.Debug("UpdateNFInstance START")
	defer logger.Debug("UpdateNFInstance END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["nfInstanceID"] = c.Param("nfInstanceID")
	logger.Debug("request:%#+v", request)

	patchData := []model.PatchItem{}
	err := c.ShouldBindJSON(&patchData)
	if err != nil {
		logger.Error("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := scenario.UpdateNFInstance(request, &patchData)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusCreated, patchData)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
}

// DeregisterNFInstance - Deregisters a given NF Instance
func DeregisterNFInstance(c *gin.Context) {

	logger.Debug("DeregisterNFInstance START")
	defer logger.Debug("DeregisterNFInstance END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["nfInstanceID"] = c.Param("nfInstanceID")
	logger.Debug("request:%#+v", request)

	// Call Scenario Function
	response, err := scenario.DeregisterNFInstance(request)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}
