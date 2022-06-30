package scenario

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
)

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
}

func GetFilterUeDataInfo(request model.Request) (bson.D, error) {

	var filter bson.D
	ueId := request.Params["ueId"]

	if ueId[0:5] == "imsi-" {
		filter = bson.D{primitive.E{Key: "ueidinfo.supi", Value: ueId}}
	} else if ueId[0:7] == "msisdn-" {
		filter = bson.D{primitive.E{Key: "ueidinfo.gpsi", Value: ueId}}
	} else if ueId[0:5] == "imei-" {
		filter = bson.D{primitive.E{Key: "ueidinfo.pei", Value: ueId}}
	} else {
		logger.Error("ueId:%v", ueId)
		return filter, fmt.Errorf(ErrorDetailDataNotFound)
	}

	return filter, nil
}

func GetUpdateUeDataInfo(request model.Request, jsonData *model.Amf3GppAccessRegistration) UeDataInfo {

	update := UeDataInfo{}
	ueId := request.Params["ueId"]

	if ueId[0:5] == "imsi-" {
		update.UeIdInfo.Supi = ueId
	} else if ueId[0:7] == "msisdn-" {
		update.UeIdInfo.Gpsi = ueId
	} else if ueId[0:5] == "imei-" {
		update.UeIdInfo.Pei = ueId
	} else {
		logger.Error("ueId:%v", ueId)
	}

	if jsonData.Pei != nil {
		update.UeIdInfo.Pei = *jsonData.Pei
	}
	if jsonData.Supi != nil {
		update.UeIdInfo.Supi = *jsonData.Supi
	}
	update.AmfAccessReg = jsonData

	return update
}
