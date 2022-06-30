package scenario

import (
	"gogc/src/common/db"
	"gogc/src/common/logger"
	"gogc/src/model"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	db.Init()
	code := m.Run()

	os.Exit(code)
}

func TestQueryAmfContext3gppOK(t *testing.T) {

	logger.Debug("TestQueryAmfContext3gppOK START")
	defer logger.Debug("TestQueryAmfContext3gppOK END")

	var request model.Request
	request.Params = map[string]string{"ueId": "imsi-11223344"}
	response, err := QueryAmfContext3gpp(request)
	logger.Debug("response:%#v, err:%v", response, err)
	if err != nil {
		t.Errorf("TestSingleOK NG")
	}
}
