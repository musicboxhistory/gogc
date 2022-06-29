package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
)

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
}

func GetFindFilter(request model.Request) (UeDataInfo, error) {

//	filter := UeDataInfo{}
	var filter UeDataInfo
	ueId := request.Params["ueId"]

	if ueId[0:5] == "imsi-" {
		filter = UeDataInfo{UeIdInfo: model.UeIdentityInfo{Supi: ueId}}
//		filter.UeIdInfo.Supi = &ueId
	} else if ueId[0:7] == "msisdn-" {
		filter = UeDataInfo{UeIdInfo: model.UeIdentityInfo{Gpsi: ueId}}
//		filter.UeIdInfo.Gpsi = &ueId
	} else if ueId[0:5] == "imei-" {
		filter = UeDataInfo{UeIdInfo: model.UeIdentityInfo{Pei: ueId}}
//		filter.UeIdInfo.Pei = &ueId
	} else {
		logger.Error("ueId:%v", ueId)
		return filter, fmt.Errorf(ErrorDetailDataNotFound)
	}

	return filter, nil
}
