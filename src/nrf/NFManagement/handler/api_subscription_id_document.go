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

// RemoveSubscription - Deletes a subscription
func RemoveSubscription(c *gin.Context) {

	logger.Debug("DeregisterNFInstance START")
	defer logger.Debug("DeregisterNFInstance END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["subscriptionID"] = c.Param("subscriptionID")
	logger.Debug("request:%#+v", request)

	// Call Scenario Function
	response, err := scenario.RemoveSubscription(request)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}

// UpdateSubscription - Updates a subscription
func UpdateSubscription(c *gin.Context) {

	logger.Debug("UpdateSubscription START")
	defer logger.Debug("UpdateSubscription END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["subscriptionID"] = c.Param("subscriptionID")
	logger.Debug("request:%#+v", request)

	patchData := []model.PatchItem{}
	err := c.ShouldBindJSON(&patchData)
	if err != nil {
		logger.Error("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := scenario.UpdateSubscription(request, &patchData)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusCreated, patchData)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
}
