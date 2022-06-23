package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func GetEquipmentStatus(input EquipmentStatus) (interface{}, error) {

	logger.Snap("GetEquipmentStatus START")
	defer logger.Snap("GetEquipmentStatus END")

	// Variable Declaration
	var response model.EirResponseData

	// Get Equipment Status
	equipmentStatus := GetStatus(input)
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
