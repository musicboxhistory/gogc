package scenario

import (
	"gogc/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFilter(request model.Request) bson.D {

	nfInstanceID := request.Params["nfInstanceID"]
	filter := bson.D{primitive.E{Key: "nfInstanceID", Value: nfInstanceID}}

	return filter
}
