package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetNFInstance(request model.Request) (interface{}, error) {

	logger.Debug("GetNFInstance START")
	defer logger.Debug("GetNFInstance END")

	// Find NF Instance
	response, err := FindNFInstance(request)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailInvalidRequest
		cause := InvalidRequest
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailInvalidRequest)
		return problemDetail, err
	}

	return response, nil
}

func FindNFInstance(request model.Request) (*model.NfProfile, error) {

	logger.Debug("FindNFInstance START")
	defer logger.Debug("FindNFInstance END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)
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
	response := model.NfProfile{}
	err = bson.Unmarshal(data, &response)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return &response, nil
}
