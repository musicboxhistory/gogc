package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Post(nfType string, input []EquipmentStatus) error {

	logger.Snap("Post EquipmentStatus Station START")
	defer logger.Snap("Post EquipmentStatus Station END")

	filter := make([]interface{}, 0)
	for _, value := range input {
		doc := bson.D{primitive.E{Key: Key, Value: value.Key}, primitive.E{Key: Status, Value: value.Status}}
		filter = append(filter, doc)
	}

	logger.Snap("filter:%#+v", filter)
	err := db.InsertMany(nfType, db.EquipmentStatus, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
