package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gogc/src/common/logger"
	"gogc/src/model"
	"os"
	"testing"
)

type TestData struct {
	A string  `json:"a" bson:"a"`
	B string  `json:"b" bson:"b"`
	C string  `json:"c,omitempty" bson:"c,omitempty"`
	D string  `json:"d,omitempty" bson:"d,omitempty"`
	E *string `json:"e" bson:"e"`
	F *string `json:"f" bson:"f"`
	G string  `json:"g,omitempty" bson:"g,omitempty"`
	H string  `json:"h,omitempty" bson:"h,omitempty"`
}

func TestMain(m *testing.M) {

	Init()
	code := m.Run()

	os.Exit(code)
}

func TestSingleOK(t *testing.T) {

	logger.Debug("TestSingleOK START")
	defer logger.Debug("TestSingleOK END")

	// InsertOne
	filter := bson.D{primitive.E{Key: "a", Value: "pei-11227788"}, primitive.E{Key: "b", Value: model.WHITELISTED}}
	result, err := InsertOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)

	// FindOne
	filter = bson.D{primitive.E{Key: "a", Value: "pei-11227788"}}
	result, err = FindOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)

	// DeleteOne
	result, err = DeleteOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)
}

func TestSingleStructOK(t *testing.T) {

	logger.Debug("TestSingleStructOK START")
	defer logger.Debug("TestSingleStructOK END")

	var response TestData

	// InsertOne
	//key := "Key test"
	//status := "Status test"
	filter := TestData{A: "pei-11227788", B: "WHITELISTED"}
	result, err := InsertOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)

	// FindOne
	result, err = FindOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)
	data, err := bson.Marshal(result)
	err = bson.Unmarshal(data, &response)
	logger.Debug("response:%#+v", response)

	// DeleteOne
	result, err = DeleteOne("test", "collection", filter)
	if err != nil {
		logger.Debug("err:%v", err)
		t.Errorf("TestSingleOK NG")
	}
	logger.Debug("result:%#+v", result)
}
