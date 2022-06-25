package scenario

import (
	"fmt"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func SearchNFInstances(request model.Request) (interface{}, error) {

	logger.Debug("GetEquipmentStatus START")
	defer logger.Debug("GetEquipmentStatus END")

	// Get SearchResult
	response := GetSearchResult(request)
	if response != nil {
		return response, nil
	}

	// Set Error Details
	status := http.StatusNotFound
	detail := ErrorDetailTargetUnknown
	cause := TargetUnknown
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailTargetUnknown)

	return problemDetail, err
}
