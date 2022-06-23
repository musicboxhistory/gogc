package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
)

func Get(nfType string) (interface{}, error) {

	logger.Snap("Get EquipmentStatus Station START")
	defer logger.Snap("Get EquipmentStatus Station END")

	filter := bson.D{}
	response, err := db.Find(nfType, db.EquipmentStatus, filter)

	return response, err
}
