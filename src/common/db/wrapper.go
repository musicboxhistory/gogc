package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() error {

	var err error
	uri := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	return nil
}

func Close() error {

	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func FindOne(database string, collection string, filter interface{}, opts ...*options.FindOneOptions) (RespData, error) {

	var result RespData

	coll := client.Database(database).Collection(collection)
	err := coll.FindOne(context.TODO(), filter, opts...).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func Find(database string, collection string, filter interface{}, opts ...*options.FindOptions) ([]RespData, error) {

	var result []RespData

	coll := client.Database(database).Collection(collection)
	cur, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
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

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func InsertMany(database string, collection string, filter []interface{}) error {

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) error {

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMany(database string, collection string, filter interface{}, opts ...*options.DeleteOptions) error {

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteMany(context.TODO(), filter, opts...)
	if err != nil {
		return err
	}

	return nil
}

func UpdataOne(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateOne(context.TODO(), filter, update, opts...)
	if err != nil {
		return err
	}

	return nil
}

func UpdataMany(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateMany(context.TODO(), filter, update, opts...)
	if err != nil {
		return err
	}

	return nil
}
