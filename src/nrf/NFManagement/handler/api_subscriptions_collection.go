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

// CreateSubscription - Create a new subscription
func CreateSubscription(c *gin.Context) {

	logger.Debug("CreateSubscription START")
	defer logger.Debug("CreateSubscription END")

	// Get Parameter
	request := signal.RequestInit(c)
	logger.Debug("request:%#+v", request)

	subscriptionData := model.SubscriptionData{}
	err := c.ShouldBindJSON(&subscriptionData)
	if err != nil {
		logger.Error("err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Call Scenario Function
	response, err := scenario.CreateSubscription(request, &subscriptionData)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusCreated, subscriptionData)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
}
