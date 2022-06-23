package common

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
)

func CheckNftype(nfType string) bool {

	switch nfType {
	case db.Database5geir:
	case db.DatabaseAf:
	case db.DatabaseAmf:
	case db.DatabaseAusf:
	case db.DatabaseLmf:
	case db.DatabaseNrf:
	case db.DatabaseNssf:
	case db.DatabasePcf:
	case db.DatabaseSmf:
	case db.DatabaseSmsf:
	case db.DatabaseUdm:
	case db.DatabaseUdr:
	case db.DatabaseUpf:
		return true
	default:
		logger.Error("nfType:%v", nfType)
	}

	return false
}
