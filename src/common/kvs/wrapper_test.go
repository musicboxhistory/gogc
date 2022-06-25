package kvs

import (
	"gogc/src/common/logger"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	Init()
	code := m.Run()

	os.Exit(code)
}

func TestSingleOK(t *testing.T) {

	logger.Debug("TestSingleOK START")
	defer logger.Debug("TestSingleOK END")

	// Set
	result, err := Set("single", 1, time.Hour*1)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%v", result)

	// Get
	result, err = Get("single")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%v", result)

	// Del
	result, err = Del("single")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%v", result)
}

func TestMultiOK(t *testing.T) {

	logger.Debug("TestMultiOK START")
	defer logger.Debug("TestMultiOK END")

	// HSet
	result, err := HSet("multi", "sample1", "TestData1")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)
	result, err = HSet("multi", "sample2", "TestData2")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)

	// HGet
	result, err = HGet("multi", "sample2")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)

	// HGetAll
	result, err = HGetAll("multi")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)

	// HDel
	result, err = HDel("multi", "sample1")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)
	result, err = HDel("multi", "sample2")
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Debug("result:%v", result)
}
