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

func DeregisterNFInstance(request model.Request) (interface{}, error) {

	logger.Debug("DeregisterNFInstance START")
	defer logger.Debug("DeregisterNFInstance END")

	// Delete NF Instance
	err := DeleteNFInstance(request)
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

func GetNFInstance(request model.Request) (interface{}, error) {

	logger.Debug("GetNFInstance START")
	defer logger.Debug("GetNFInstance END")

	// Find NF Instance
	response, err := FindNFInstance(request)
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

func RegisterNFInstance(request model.Request, nfProfile *model.NfProfile) (interface{}, error) {

	logger.Debug("RegisterNFInstance START")
	defer logger.Debug("RegisterNFInstance END")

	// Update NF Instance
	err := PutNFInstance(request, nfProfile)
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

func UpdateNFInstance(request model.Request, nfProfile *model.NfProfile) (interface{}, error) {

	logger.Debug("UpdateNFInstance START")
	defer logger.Debug("UpdateNFInstance END")

	// Patch NF Instance
	response, err := PatchNFInstance(request, nfProfile)
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

func PutNFInstance(request model.Request, nfProfile *model.NfProfile) error {

	logger.Debug("PutNFInstance START")
	defer logger.Debug("PutNFInstance END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)

	if err == nil {
		// Update DB
		logger.Debug("Update DB")
		_, err = db.UpdataOne(db.DatabaseNrf, db.NFProfile, filter, nfProfile)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		_, err = db.InsertOne(db.DatabaseNrf, db.NFProfile, nfProfile)
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

func PatchNFInstance(request model.Request, nfProfile *model.NfProfile) (*model.NfProfile, error) {

	logger.Debug("PatchNFInstance START")
	defer logger.Debug("PatchNFInstance END")

	return nil, nil
}

func DeleteNFInstance(request model.Request) error {

	logger.Debug("DeleteNFInstance START")
	defer logger.Debug("DeleteNFInstance END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)

	if err != nil {
		logger.Debug("err:%v", err)
		return err
	}

	// Delete DB
	logger.Debug("Update DB")
	_, err = db.DeleteOne(db.DatabaseNrf, db.NFProfile, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
