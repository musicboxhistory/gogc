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
	"gogc/src/common/logger"
	"gogc/src/common/signal"
	"gogc/src/udm/SubscriberDataManagement/scenario"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAmData - retrieve a UE's Access and Mobility Subscription Data
func GetAmData(c *gin.Context) {

	logger.Debug("GetAmData START")
	defer logger.Debug("GetAmData END")

	// Get Parameter
	request := signal.RequestInit(c)
	request.Params["supi"] = c.Param("supi")
	logger.Debug("request:%#+v", request)

	// Call Scenario Function
	response, err := scenario.GetAmData(request)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}
