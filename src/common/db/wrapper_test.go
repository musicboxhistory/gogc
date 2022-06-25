package db

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

	logger.Snap("TestSingleOK START")
	defer logger.Snap("TestSingleOK END")

	// Set
	result, err := Set("single", 1, time.Hour*1)
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Snap("result:%v", result)

	// Get
	result, err = Get("single")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Snap("result:%v", result)

	// Del
	result, err = Del("single")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Snap("result:%v", result)
}

func TestMultiOK(t *testing.T) {

	logger.Snap("TestMultiOK START")
	defer logger.Snap("TestMultiOK END")

	// HSet
	result, err := HSet("multi", "sample1", "TestData1")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)
	result, err = HSet("multi", "sample2", "TestData2")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)

	// HGet
	result, err = HGet("multi", "sample2")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)

	// HGetAll
	result, err = HGetAll("multi")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)

	// HDel
	result, err = HDel("multi", "sample1")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)
	result, err = HDel("multi", "sample2")
	if err != nil {
		logger.Snap("err:%v", err)
		t.Errorf("TestMultiOK NG")
	}
	logger.Snap("result:%v", result)
}
