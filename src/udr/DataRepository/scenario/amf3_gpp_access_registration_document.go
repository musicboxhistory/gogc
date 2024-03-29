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

func AmfContext3gpp(request model.Request) (interface{}, error) {

	logger.Debug("AmfContext3gpp START")
	defer logger.Debug("AmfContext3gpp END")

	// To modify the AMF context data of a UE using 3gpp access in the UDR

	// Set Error Details
	status := http.StatusNotFound
	detail := ErrorDetailUserNotFoud
	cause := UserNotFoud
	problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
	err := fmt.Errorf(ErrorDetailUserNotFoud)

	return problemDetail, err
}

func CreateAmfContext3gpp(request model.Request, jsonData *model.Amf3GppAccessRegistration) (interface{}, error) {

	logger.Debug("CreateAmfContext3gpp START")
	defer logger.Debug("CreateAmfContext3gpp END")

	// To store the AMF context data of a UE using 3gpp access in the UDR
	err := PutAmfContext3gpp(request, jsonData)
	if err != nil {
		// Set Error Details
		status := http.StatusNotFound
		detail := ErrorDetailUserNotFoud
		cause := UserNotFoud
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
		return problemDetail, err
	}

	return nil, nil
}

func QueryAmfContext3gpp(request model.Request) (interface{}, error) {

	logger.Debug("QueryAmfContext3gpp START")
	defer logger.Debug("QueryAmfContext3gpp END")

	// Retrieves the AMF context data of a UE using 3gpp access
	response, err := GetAmfContext3gpp(request)
	if err != nil {
		// Set Error Details
		status := http.StatusNotFound
		detail := ErrorDetailUserNotFoud
		cause := UserNotFoud
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}
		err := fmt.Errorf(ErrorDetailUserNotFoud)

		return problemDetail, err
	}

	return response, nil
}

func PutAmfContext3gpp(request model.Request, jsonData *model.Amf3GppAccessRegistration) error {

	logger.Debug("CreateAmfContext3gpp START")
	defer logger.Debug("CreateAmfContext3gpp END")

	// Find DB
	filter, err := GetFilterUeDataInfo(request)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(db.DatabaseUdr, db.UeDataInfo, filter)

	if err == nil {
		// Update DB
		logger.Debug("Update DB")
		// Bson Marshal
		data, err := bson.Marshal(result)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
		update := UeDataInfo{}
		err = bson.Unmarshal(data, &update)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
		update.AmfAccessReg = jsonData
		_, err = db.UpdataOne(db.DatabaseUdr, db.UeDataInfo, filter, update)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		update := GetUpdateUeDataInfo(request)
		update.AmfAccessReg = jsonData
		_, err = db.InsertOne(db.DatabaseUdr, db.UeDataInfo, update)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else {
		return err
	}

	return nil
}

func GetAmfContext3gpp(request model.Request) (*model.Amf3GppAccessRegistration, error) {

	logger.Debug("GetAmfContext3gpp START")
	defer logger.Debug("GetAmfContext3gpp END")

	// Set DB Filter
	filter, err := GetFilterUeDataInfo(request)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}
	logger.Debug("filter:%#v", filter)

	// Find DB
	result, err := db.FindOne(db.DatabaseUdr, db.UeDataInfo, filter)
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
	response := UeDataInfo{}
	err = bson.Unmarshal(data, &response)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return response.AmfAccessReg, nil
}
