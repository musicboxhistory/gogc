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
	detail := ErrorDetailUserNotFound
	cause := UserNotFound
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailUserNotFound)

	return problemDetail, err
}
