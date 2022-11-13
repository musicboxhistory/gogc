package scenario

import (
	"encoding/json"
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	jsonpatch "github.com/evanphx/json-patch"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSubscription(request model.Request, patchData *[]model.PatchItem) (interface{}, error) {

	logger.Debug("UpdateSubscription START")
	defer logger.Debug("UpdateSubscription END")

	// Patch NF Instance
	response, err := PatchSubscription(request, patchData)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailInvalidRequest
		cause := InvalidRequest
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailInvalidRequest)
		return problemDetail, err
	}

	return response, nil
}

func PatchSubscription(request model.Request, patchData *[]model.PatchItem) (*model.SubscriptionData, error) {

	logger.Debug("PatchSubscription START")
	defer logger.Debug("PatchSubscription END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	result, err := db.FindOne(db.DatabaseNrf, db.SubscriptionData, filter)

	if err != nil {
		logger.Debug("err:%v", err)
		return nil, err
	}

	// Bson Marshal
	original, err := bson.Marshal(result)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	// Json Patch
	patchJSON, err := json.Marshal(patchData)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	modified, err := patch.Apply(original)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	var subscriptionData model.SubscriptionData
	err = json.Unmarshal(modified, &subscriptionData)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	// Update DB
	logger.Debug("Update DB")
	_, err = db.UpdataOne(db.DatabaseNrf, db.SubscriptionData, filter, subscriptionData)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return &subscriptionData, nil
}
