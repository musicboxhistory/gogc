package scenario

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"gogc/src/station/equipmentstatus"
	"sync"
	"time"
)

// global Variable
var equipmentStatusList []equipmentstatus.EquipmentStatus
var mutex sync.Mutex

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
	go GetDatabase()
}

func GetDatabase() {

	logger.Debug("GetDatabase START")
	defer logger.Debug("GetDatabase END")

	filter := bson.D{{}}
	for {
		// Get Database
		result, err := db.Find(db.Database5geir, db.EquipmentStatus, filter)
		if err != nil {
			logger.Error("err:%v", err)
			time.Sleep(time.Second * 1)
			continue
		}

		// Unmarshal Database
		list := make([]equipmentstatus.EquipmentStatus, len(result))
		for idx := 0; idx < len(result); idx++ {
			data, err := bson.Marshal(result[idx])
			if err != nil {
				logger.Error("err:%v", err)
				time.Sleep(time.Second * 1)
				continue
			}
			err = bson.Unmarshal(data, &list[idx])
			if err != nil {
				logger.Error("err:%v", err)
				time.Sleep(time.Second * 1)
				continue
			}
		}

		// Set Database From Variable
		mutex.Lock()
		equipmentStatusList = list
		mutex.Unlock()

		// sleep
		time.Sleep(time.Second * 1)
	}
}

func GetStatus(request model.Request) model.EquipmentStatus {

	logger.Debug("GetStatus START")
	defer logger.Debug("GetStatus END")

	mutex.Lock()
	defer mutex.Unlock()

	for _, value := range equipmentStatusList {

		// Check Pei
		if request.Query["pei"] != nil && value.Key == fmt.Sprintf("pei-%s", request.Query["pei"][0]) {
			return value.Status
		}

		// Check Supi
		if request.Query["supi"] != nil && value.Key == fmt.Sprintf("supi-%s", request.Query["supi"][0]) {
			return value.Status
		}

		// Check Gpsi
		if request.Query["gpsi"] != nil && value.Key == fmt.Sprintf("gpsi-%s", request.Query["gpsi"][0]) {
			return value.Status
		}
	}

	return ""
}
