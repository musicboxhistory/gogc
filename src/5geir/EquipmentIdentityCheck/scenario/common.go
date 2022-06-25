package scenario

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"sync"
	"time"
)

// global Variable
var equipmentStatusList map[string]model.EquipmentStatus
var mutex sync.Mutex

func Init() {

	logger.Snap("Init START")
	defer logger.Snap("Init END")

	db.Init()
	go GetDatabase()
}

func GetDatabase() {

	filter := bson.D{{}}
	for {
		// Get Database
		_, err := db.Find(db.Database5geir, db.EquipmentStatus, filter)
		if err != nil {
			logger.Error("err: %v", err)
			time.Sleep(time.Second * 1)
			continue
		}

		// Set Database From Variable
		mutex.Lock()
		mutex.Unlock()

		// sleep
		time.Sleep(time.Second * 1)
	}
}

func GetStatus(request EquipmentStatus) model.EquipmentStatus {

	mutex.Lock()
	defer mutex.Unlock()

	// Check Pei
	equipmentStatus, ok := equipmentStatusList[fmt.Sprintf("pei-%s", request.Pei)]
	if ok {
		return equipmentStatus
	}

	// Check Supi
	equipmentStatus, ok = equipmentStatusList[fmt.Sprintf("supi-%s", request.Supi)]
	if ok {
		return equipmentStatus
	}

	// Check Gpsi
	equipmentStatus, ok = equipmentStatusList[fmt.Sprintf("gpsi-%s", request.Gpsi)]
	if ok {
		return equipmentStatus
	}

	return ""
}
