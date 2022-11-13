package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSubscription(request model.Request, subscriptionData *model.SubscriptionData) (interface{}, error) {

	logger.Debug("CreateSubscription START")
	defer logger.Debug("CreateSubscription END")

	// Update NF Instance
	err := PostSubscription(request, subscriptionData)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailInvalidRequest
		cause := InvalidRequest
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailInvalidRequest)
		return problemDetail, err
	}

	return subscriptionData, nil
}
func PostSubscription(request model.Request, subscriptionData *model.SubscriptionData) error {

	logger.Debug("PostSubscription START")
	defer logger.Debug("PostSubscription END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.SubscriptionData, filter)

	if err == nil {
		// Update DB
		logger.Debug("Update DB")
		_, err = db.UpdataOne(db.DatabaseNrf, db.SubscriptionData, filter, subscriptionData)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		_, err = db.InsertOne(db.DatabaseNrf, db.SubscriptionData, subscriptionData)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
