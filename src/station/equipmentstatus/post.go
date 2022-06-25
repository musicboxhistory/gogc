package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
)

func Post(nfType string, request []EquipmentStatus) error {

	logger.Debug("Post EquipmentStatus Station START")
	defer logger.Debug("Post EquipmentStatus Station END")

	filter := make([]interface{}, 0)
	for _, value := range request {
		filter = append(filter, value)
	}

	logger.Debug("filter:%#+v", filter)
	_, err := db.InsertMany(nfType, db.EquipmentStatus, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
