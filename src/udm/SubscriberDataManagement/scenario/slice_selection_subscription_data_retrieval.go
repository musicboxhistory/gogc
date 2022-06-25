package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func GetNSSAI(request model.Request) (interface{}, error) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Get SearchResult

	// Set Error Details
	status := http.StatusNotFound
	detail := ErrorDetailQueryUnknown
	cause := NssaiUnknown
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailQueryUnknown)

	return problemDetail, err
}
