package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func SearchNFInstances(request model.Request) (interface{}, error) {

	logger.Debug("SearchNFInstances START")
	defer logger.Debug("SearchNFInstances END")

	// Find NF Instance
	response, err := FindNFInstances(request)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailTargetUnknown
		cause := TargetUnknown
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailTargetUnknown)
		return problemDetail, err
	}

	return response, nil
}

func FindNFInstances(request model.Request) (*model.SearchResult, error) {

	logger.Debug("FindNFInstances START")
	defer logger.Debug("FindNFInstances END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	result, err := db.Find(db.DatabaseNrf, db.NFProfile, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	// Bson Marshal
	data, err := bson.Marshal(result)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}
	nfProfiles := []model.NfProfile{}
	err = bson.Unmarshal(data, &nfProfiles)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	response := model.SearchResult{}
	response.NfInstances = &nfProfiles
	validityPeriod := int32(30)
	response.ValidityPeriod = &validityPeriod

	return &response, nil
}
