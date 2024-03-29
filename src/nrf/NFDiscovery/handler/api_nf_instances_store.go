/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handler

import (
	"gogc/src/common/logger"
	"gogc/src/common/signal"
	"gogc/src/model"
	"gogc/src/nrf/NFDiscovery/scenario"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchNFInstances - Search a collection of NF Instances
func SearchNFInstances(c *gin.Context) {

	logger.Debug("SearchNFInstances START")
	defer logger.Debug("SearchNFInstances END")

	// Get Parameter
	request := signal.RequestInit(c)
	logger.Debug("request:%#+v", request)

	// Check Mandatory Parameter
	if request.Query["target-nf-type"] == nil || request.Query["requester-nf-type"] == nil {
		status := int32(http.StatusBadRequest)
		detail := scenario.ErrorDetailMandatoryIeIncorrect
		cause := scenario.MandatoryIeIncorrect
		response := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call Scenario Function
	response, err := scenario.SearchNFInstances(request)
	// logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}
