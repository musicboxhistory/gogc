package scenario

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
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
	filter, err := GetFindFilter(request)
	_, err = db.FindOne(db.DatabaseUdr, db.UeDataInfo, filter)

	update := filter
	if jsonData.Pei != nil {
		update.UeIdInfo.Pei = jsonData.Pei
	}
	if jsonData.Supi != nil {
		update.UeIdInfo.Supi = jsonData.Supi
	}
	update.AmfAccessReg = jsonData
	if err == nil {
		// Update DB
		_, err = db.UpdataOne(db.DatabaseUdr, db.UeDataInfo, filter, update)
		if err != nil {
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		_, err = db.InsertOne(db.DatabaseUdr, db.UeDataInfo, update)
		if err != nil {
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
	filter, err := GetFindFilter(request)
	if err != nil {
		return nil, err
	}

	// Find DB
	result, err := db.FindOne(db.DatabaseUdr, db.UeDataInfo, filter)
	if err != nil {
		return nil, err
	}

	// Bson Marshal
	data, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}
	response := UeDataInfo{}
	err = bson.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.AmfAccessReg, nil
}
