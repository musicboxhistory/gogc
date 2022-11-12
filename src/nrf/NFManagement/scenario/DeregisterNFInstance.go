package scenario

import (
	"fmt"
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"net/http"
)

func DeregisterNFInstance(request model.Request) (interface{}, error) {

	logger.Debug("DeregisterNFInstance START")
	defer logger.Debug("DeregisterNFInstance END")

	// Delete NF Instance
	err := DeleteNFInstance(request)
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

func DeleteNFInstance(request model.Request) error {

	logger.Debug("DeleteNFInstance START")
	defer logger.Debug("DeleteNFInstance END")

	// Find DB
	filter := GetFilter(request)
	logger.Debug("filter:%#v", filter)
	_, err := db.FindOne(db.DatabaseNrf, db.NFProfile, filter)

	if err != nil {
		logger.Debug("err:%v", err)
		return err
	}

	// Delete DB
	logger.Debug("Delete DB")
	_, err = db.DeleteOne(db.DatabaseNrf, db.NFProfile, filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
