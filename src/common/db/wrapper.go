package db

import (
	"context"
	"fmt"
	"gogc/src/common/logger"

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

func FindOne(database string, collection string, filter interface{}, opts ...*options.FindOneOptions) (interface{}, error) {

	var result interface{}

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	err := coll.FindOne(context.TODO(), filter, opts...).Decode(&result)
	if err != nil {
		logger.Error("err:%v", err)
		return result, err
	}

	return result, nil
}

func Find(database string, collection string, filter interface{}, opts ...*options.FindOptions) ([]interface{}, error) {

	var result []interface{}

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
		var dec RespData
		cur.Decode(&dec)
		result = append(result, dec)
	}

	return result, nil
}

func InsertOne(database string, collection string, filter interface{}) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertOne(context.TODO(), filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func InsertMany(database string, collection string, filter []interface{}) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertMany(context.TODO(), filter)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func DeleteOne(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func DeleteMany(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteMany(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func UpdataOne(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateOne(context.TODO(), filter, update, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}

func UpdataMany(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {

	if client == nil {
		return fmt.Errorf("client nil")
	}

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateMany(context.TODO(), filter, update, opts...)
	if err != nil {
		logger.Error("err:%v", err)
		return err
	}

	return nil
}
