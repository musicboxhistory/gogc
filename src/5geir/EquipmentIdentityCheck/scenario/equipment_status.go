package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func GetEquipmentStatus(request model.Request) (interface{}, error) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Get Equipment Status
	response := GetStatus(request)
	if response.Status == "" {
		// Set Error Details
		status := http.StatusNotFound
		detail := ErrorDetailEquipmentUnknown
		cause := EquipmentUnknown
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailEquipmentUnknown)
		return problemDetail, err
	}

	return response, nil
}

func GetStatus(request model.Request) model.EirResponseData {

	logger.Debug("GetStatus START")
	defer logger.Debug("GetStatus END")

	var response model.EirResponseData

	mutex.Lock()
	defer mutex.Unlock()

	for _, value := range equipmentStatusList {

		// Check Query
		if request.Query["pei"] != nil && value.Key == request.Query["pei"][0] ||
			request.Query["supi"] != nil && value.Key == request.Query["supi"][0] ||
			request.Query["gpsi"] != nil && value.Key == request.Query["gpsi"][0] {
			response.Status = value.Status
		}
	}

	return response
}
