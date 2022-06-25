package common

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
)

func CheckNftype(nfType string) bool {

	switch nfType {
	case db.Database5geir:
		fallthrough
	case db.DatabaseAf:
		fallthrough
	case db.DatabaseAmf:
		fallthrough
	case db.DatabaseAusf:
		fallthrough
	case db.DatabaseLmf:
		fallthrough
	case db.DatabaseNrf:
		fallthrough
	case db.DatabaseNssf:
		fallthrough
	case db.DatabasePcf:
		fallthrough
	case db.DatabaseSmf:
		fallthrough
	case db.DatabaseSmsf:
		fallthrough
	case db.DatabaseUdm:
		fallthrough
	case db.DatabaseUdr:
		fallthrough
	case db.DatabaseUpf:
		logger.Debug("Check NF Type OK")
		return true
	default:
		logger.Error("nfType:%v", nfType)
	}

	return false
}
