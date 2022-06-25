package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func GetEquipmentStatus(request EirRequestData) (interface{}, error) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Variable Declaration
	var response model.EirResponseData

	// Get Equipment Status
	equipmentStatus := GetStatus(request)
	if equipmentStatus != "" {
		response.Status = equipmentStatus
		return response, nil
	}

	// Set Error Details
	status := http.StatusNotFound
	detail := ErrorDetailEquipmentUnknown
	cause := EquipmentUnknown
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailEquipmentUnknown)

	return problemDetail, err
}
