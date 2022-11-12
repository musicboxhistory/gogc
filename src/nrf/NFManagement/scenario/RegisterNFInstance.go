package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterNFInstance(request model.Request, nfProfile *model.NfProfile) (interface{}, error) {

	logger.Debug("RegisterNFInstance START")
	defer logger.Debug("RegisterNFInstance END")

	// Update NF Instance
	err := PutNFInstance(request, nfProfile)
	if err != nil {
		// Set Error Details
		status := int32(http.StatusNotFound)
		detail := ErrorDetailInvalidRequest
		cause := InvalidRequest
		problemDetail := model.ProblemDetails{Status: &status, Detail: &detail, Cause: &cause}

		err := fmt.Errorf(ErrorDetailInvalidRequest)
		return problemDetail, err
	}

	return nfProfile, nil
}
func PutNFInstance(request model.Request, nfProfile *model.NfProfile) error {

	logger.Debug("PutNFInstance START")
	defer logger.Debug("PutNFInstance END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)

	if err == nil {
		// Update DB
		logger.Debug("Update DB")
		_, err = db.UpdataOne(db.DatabaseNrf, db.NFProfile, filter, nfProfile)
		if err != nil {
			logger.Error("err:%v", err)
			return err
		}
	} else if err == mongo.ErrNoDocuments {
		// Insert DB
		logger.Debug("Insert DB")
		_, err = db.InsertOne(db.DatabaseNrf, db.NFProfile, nfProfile)
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
