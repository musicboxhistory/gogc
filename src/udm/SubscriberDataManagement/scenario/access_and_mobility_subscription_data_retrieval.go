package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func GetAmData(request model.Request) (interface{}, error) {

	logger.Debug("GetAmData START")
	defer logger.Debug("GetAmData END")

	// Get SearchResult

	// Set Error Details
	status := http.StatusNotFound
	detail := ErrorDetailUserNotFound
	cause := UserNotFound
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailUserNotFound)

	return problemDetail, err
}
