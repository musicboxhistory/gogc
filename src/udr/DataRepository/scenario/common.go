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

	filter := UeDataInfo{}
	ueId := request.Params["ueId"]

	if ueId[0:4] == "imsi-" {
		filter.UeIdInfo.Supi = &ueId
	} else if ueId[0:6] == "msisdn-" {
		filter.UeIdInfo.Gpsi = &ueId
	} else if ueId[0:4] == "imei-" {
		filter.UeIdInfo.Pei = &ueId
	} else {
		return filter, fmt.Errorf(ErrorDetailDataNotFound)
	}

	return filter, nil
}
