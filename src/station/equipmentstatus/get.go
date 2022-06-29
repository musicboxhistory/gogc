package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetOne(nfType string, key string) (EquipmentStatus, error) {

	logger.Debug("GetOne EquipmentStatus Station START")
	defer logger.Debug("GetOne EquipmentStatus Station END")

	var response EquipmentStatus

	filter := EquipmentStatus{Key: key}
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(nfType, db.EquipmentStatus, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Debug("err:%v", err)
		} else {
			logger.Error("err:%v", err)
		}
		return response, err
	}

	data, err := bson.Marshal(result)
	if err != nil {
		logger.Error("err:%v", err)
		return response, err
	}
	err = bson.Unmarshal(data, &response)
	if err != nil {
		logger.Error("err:%v", err)
		return response, err
	}

	return response, nil
}

func Get(nfType string) ([]EquipmentStatus, error) {

	logger.Debug("Get EquipmentStatus Station START")
	defer logger.Debug("Get EquipmentStatus Station END")

	filter := bson.D{}
	result, err := db.Find(nfType, db.EquipmentStatus, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	response := make([]EquipmentStatus, len(result))
	for idx := 0; idx < len(result); idx++ {
		data, err := bson.Marshal(result[idx])
		if err != nil {
			logger.Error("err:%v", err)
			return nil, err
		}
		err = bson.Unmarshal(data, &response[idx])
		if err != nil {
			logger.Error("err:%v", err)
			return nil, err
		}
	}

	return response, nil
}
