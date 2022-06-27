package db

import (
	"context"
	"fmt"
	"gogc/src/common/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() error {

	var err error
	uri := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func Close() error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func FindOne(database string, collection string, filter interface{}, opts ...*options.FindOneOptions) (bson.D, error) {

	var dec interface{}
	var result bson.D

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	err := coll.FindOne(context.TODO(), filter, opts...).Decode(&dec)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			logger.Error("err:%v", err)
		}
		return nil, err
	}
	bsonData := dec.(bson.D)
	for i, value := range bsonData {
		if i == 0 {
			continue
		}
		result = append(result, value)
	}

	return result, nil
}

func Find(database string, collection string, filter interface{}, opts ...*options.FindOptions) ([]bson.D, error) {

	var result []bson.D

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	cur, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return result, err
	}

	for cur.Next(context.TODO()) {
		var dec interface{}
		var single bson.D
		cur.Decode(&dec)
		bsonData := dec.(bson.D)
		for i, value := range bsonData {
			if i == 0 {
				continue
			}
			single = append(single, value)
		}
		result = append(result, single)
	}

	return result, nil
}

func InsertOne(database string, collection string, filter interface{}) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.InsertOne(context.TODO(), filter)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}

func InsertMany(database string, collection string, filter []interface{}) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.InsertMany(context.TODO(), filter)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}

func DeleteOne(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}

func DeleteMany(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.DeleteMany(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}

func UpdataOne(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.UpdateOne(context.TODO(), filter, update, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}

func UpdataMany(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	result, err := coll.UpdateMany(context.TODO(), filter, update, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return result, nil
}
