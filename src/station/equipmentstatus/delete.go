package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(nfType string, key string) error {

	logger.Snap("Delete EquipmentStatus Station START")
	defer logger.Snap("Delete EquipmentStatus Station END")

	filter := bson.D{primitive.E{Key: Key, Value: key}}
	err := db.DeleteOne(nfType, db.EquipmentStatus, filter)

	return err
}
