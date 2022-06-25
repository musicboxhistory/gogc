package scenario

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
)

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
}
