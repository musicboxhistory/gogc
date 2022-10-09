package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NSSelectionGet(request model.Request) (interface{}, error) {

	logger.Debug("NSSAIAvailabilityDelete START")
	defer logger.Debug("NSSAIAvailabilityDelete END")

	// Delete NF Instance
	err := FindNSSelection(request)
	if err != nil {
		// Set Error Details
		status := http.StatusNotFound
		detail := ErrorDetailEquipmentUnknown
		cause := EquipmentUnknown
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailEquipmentUnknown)
		return problemDetail, err
	}

	return model.Request{}, nil
}


func FindNSSelection(request model.Request) (*model.NfProfile, error) {

	logger.Debug("PatchNSSAIAvailability START")
	defer logger.Debug("PatchNSSAIAvailability END")

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
