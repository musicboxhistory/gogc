package scenario

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOrReplaceAccessAndMobilityData(request model.Request, jsonData *model.AccessAndMobilityData) (interface{}, error) {

	logger.Debug("CreateAmfContext3gpp START")
	defer logger.Debug("CreateAmfContext3gpp END")

	// To store the AMF context data of a UE using 3gpp access in the UDR
	err := PutAccessAndMobilityData(request, jsonData)
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

func PutAccessAndMobilityData(request model.Request, jsonData *model.AccessAndMobilityData) error {

	logger.Debug("PutAccessAndMobilityData START")
	defer logger.Debug("PutAccessAndMobilityData END")

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
		update.AccessMobilityData = jsonData
		_, err = db.UpdataOne(db.DatabaseUdr, db.UeDataInfo, filter, update)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		update := GetUpdateUeDataInfo(request)
		update.AccessMobilityData = jsonData
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
