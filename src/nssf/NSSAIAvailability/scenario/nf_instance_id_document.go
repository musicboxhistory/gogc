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

func NSSAIAvailabilityDelete(request model.Request) (interface{}, error) {

	logger.Debug("NSSAIAvailabilityDelete START")
	defer logger.Debug("NSSAIAvailabilityDelete END")

	// Delete NF Instance
	err := DeleteNSSAIAvailability(request)
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

func NSSAIAvailabilityPatch(request model.Request) (interface{}, error) {

	logger.Debug("NSSAIAvailabilityPatch START")
	defer logger.Debug("NSSAIAvailabilityPatch END")

	// Find NF Instance
	response, err := PatchNSSAIAvailability(request)
	if err != nil {
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

func NSSAIAvailabilityPut(request model.Request, nfProfile *model.NfProfile) (interface{}, error) {

	logger.Debug("NSSAIAvailabilityPut START")
	defer logger.Debug("NSSAIAvailabilityPut END")

	// Update NF Instance
	err := PutNSSAIAvailability(request, nfProfile)
	if err != nil {
		// Set Error Details
		status := http.StatusNotFound
		detail := ErrorDetailEquipmentUnknown
		cause := EquipmentUnknown
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailEquipmentUnknown)
		return problemDetail, err
	}

	return nfProfile, nil
}

func DeleteNSSAIAvailability(request model.Request) (*model.NfProfile, error) {

	logger.Debug("DeleteNSSAIAvailability START")
	defer logger.Debug("DeleteNSSAIAvailability END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(db.DatabaseNssf, db.NFProfile, filter)
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

func PatchNSSAIAvailability(request model.Request) (*model.NfProfile, error) {

	logger.Debug("PatchNSSAIAvailability START")
	defer logger.Debug("PatchNSSAIAvailability END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(db.DatabaseNssf, db.NFProfile, filter)
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

func PutNSSAIAvailability(request model.Request, nfProfile *model.NfProfile) error {

	logger.Debug("PutNSSAIAvailability START")
	defer logger.Debug("PutNSSAIAvailability END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)

	if err == nil {
		// Update DB
		logger.Debug("Update DB")
		_, err = db.UpdataOne(db.DatabaseNssf, db.NFProfile, filter, nfProfile)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		_, err = db.InsertOne(db.DatabaseNssf, db.NFProfile, nfProfile)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
