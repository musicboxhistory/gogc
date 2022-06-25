package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOne(nfType string, key string) (interface{}, error) {

	logger.Snap("GetOne EquipmentStatus Station START")
	defer logger.Snap("GetOne EquipmentStatus Station END")

	filter := bson.D{primitive.E{Key: Key, Value: key}}
	response, err := db.FindOne(nfType, db.EquipmentStatus, filter)

	return response, err
}

func Get(nfType string) ([]interface{}, error) {

	logger.Snap("Get EquipmentStatus Station START")
	defer logger.Snap("Get EquipmentStatus Station END")

	filter := bson.D{}
	response, err := db.Find(nfType, db.EquipmentStatus, filter)

	return response, err
}
