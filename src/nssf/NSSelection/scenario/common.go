package scenario

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
}

func GetFilter(request model.Request) bson.D {

	filter := bson.D{}

	// Set NF Type
	if request.Query["target-nf-type"] != nil {
		filter = append(filter, primitive.E{Key: "nfType", Value: request.Query["target-nf-type"]})
	}

	return filter
}
