package scenario

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"strings"
	"time"
)

func Init() {

	logger.Debug("Init START")
	defer logger.Debug("Init END")

	db.Init()
}

func GetSearchResult(request model.Request) *model.SearchResult {

	logger.Debug("GetStatus START")
	defer logger.Debug("GetStatus END")

	response := model.SearchResult{}

	// Set validityPeriod
	response.ValidityPeriod = 3000

	// Set nfInstances
	var fqdn string
	nfProfile := model.NfProfile{}
	nfProfile.NfStatus = model.REGISTERED
	timeNow := time.Now()
	nfProfile.LoadTimeStamp = &timeNow
	targetNfType := model.NfType(strings.ToUpper(request.Query["target-nf-type"][0]))
	switch targetNfType {
	case model.EIR:
		nfProfile.NfInstanceId = "1111"
		nfProfile.NfType = model.EIR
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.AF:
		nfProfile.NfInstanceId = "2222"
		nfProfile.NfType = model.AF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.AMF:
		nfProfile.NfInstanceId = "3333"
		nfProfile.NfType = model.AMF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.AUSF:
		nfProfile.NfInstanceId = "4444"
		nfProfile.NfType = model.AUSF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.LMF:
		nfProfile.NfInstanceId = "5555"
		nfProfile.NfType = model.LMF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.NRF:
		nfProfile.NfInstanceId = "6666"
		nfProfile.NfType = model.NRF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.NSSF:
		nfProfile.NfInstanceId = "7777"
		nfProfile.NfType = model.NSSF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.PCF:
		nfProfile.NfInstanceId = "8888"
		nfProfile.NfType = model.PCF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.SMF:
		nfProfile.NfInstanceId = "9999"
		nfProfile.NfType = model.SMF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.SMSF:
		nfProfile.NfInstanceId = "AAAA"
		nfProfile.NfType = model.SMSF
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.UDM:
		nfProfile.NfInstanceId = "BBBB"
		nfProfile.NfType = model.UDM
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	case model.UDR:
		nfProfile.NfInstanceId = "CCCC"
		nfProfile.NfType = model.UDR
		fqdn = "localhost:8081"
		nfProfile.Fqdn = &fqdn
	default:
		return nil
	}
	response.NfInstances = append(response.NfInstances, nfProfile)

	return &response
}
