package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func RemoveSubscription(request model.Request) (interface{}, error) {

	logger.Debug("RemoveSubscription START")
	defer logger.Debug("RemoveSubscription END")

	// Delete NF Instance
	err := DeleteSubscription(request)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailInvalidRequest
		cause := InvalidRequest
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailInvalidRequest)
		return problemDetail, err
	}

	return model.Request{}, nil
}

func DeleteSubscription(request model.Request) error {

	logger.Debug("DeleteSubscription START")
	defer logger.Debug("DeleteSubscription END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.SubscriptionData, filter)

	if err != nil {
		logger.Debug("err:%v", err)
		return err
	}

	// Delete DB
	logger.Debug("Delete DB")
	_, err = db.DeleteOne(db.DatabaseNrf, db.SubscriptionData, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
