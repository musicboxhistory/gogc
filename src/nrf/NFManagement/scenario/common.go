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

	var filter bson.D

	if request.Params["nfInstanceID"] != "" {
		nfInstanceID := request.Params["nfInstanceID"]
		filter = bson.D{primitive.E{Key: "nfInstanceID", Value: nfInstanceID}}
	} else if request.Params["subscriptionID"] != "" {
		subscriptionID := request.Params["subscriptionID"]
		filter = bson.D{primitive.E{Key: "subscriptionID", Value: subscriptionID}}
	}

	return filter
}
