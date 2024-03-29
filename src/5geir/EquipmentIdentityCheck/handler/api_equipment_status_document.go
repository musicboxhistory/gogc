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
	"gogc/src/5geir/EquipmentIdentityCheck/scenario"
	"gogc/src/common/logger"
	"gogc/src/common/signal"
	"gogc/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEquipmentStatus - Retrieves the status of the UE
func GetEquipmentStatus(c *gin.Context) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Get Parameter
	request := signal.RequestInit(c)
	logger.Debug("request:%#+v", request)

	// Check Mandatory Parameter
	if request.Query["pei"] == nil {
		status := int32(http.StatusBadRequest)
		detail := scenario.ErrorDetailMandatoryIeIncorrect
		cause := scenario.MandatoryIeIncorrect
		response := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call Scenario Function
	response, err := scenario.GetEquipmentStatus(request)
	logger.Debug("response:%#+v, err:%v", response, err)

	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotFound, response)
	}
}
