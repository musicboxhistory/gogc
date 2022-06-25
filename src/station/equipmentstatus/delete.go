package equipmentstatus

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
)

func Delete(nfType string, key string) error {

	logger.Debug("Delete EquipmentStatus Station START")
	defer logger.Debug("Delete EquipmentStatus Station END")

	filter := EquipmentStatus{Key: key}
	_, err := db.DeleteOne(nfType, db.EquipmentStatus, filter)

	return err
}
