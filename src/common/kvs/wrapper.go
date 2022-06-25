package kvs

import (
	"fmt"
	"gogc/src/common/logger"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() error {

	options := redis.Options{Addr: "localhost:6379", DB: 0}
	client = redis.NewClient(&options)

	if client == nil {
		logger.Error("Init Error")
		return fmt.Errorf("Init Error")
	}

	return nil
}

func Close() error {

	return nil
}

func Set(key string, value interface{}, expiration time.Duration) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.Set(key, value, expiration)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func HSet(key string, filed string, value interface{}) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.HSet(key, filed, value)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func SetNX(key string, value interface{}, expiration time.Duration) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.SetNX(key, value, expiration)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func HSetNX(key string, filed string, value interface{}) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.HSetNX(key, filed, value)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func Get(key string) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.Get(key)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func HGet(key string, filed string) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.HGet(key, filed)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func HGetAll(key string) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.HGetAll(key)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func Del(key string) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.Del(key)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}

func HDel(key string, filed string) (interface{}, error) {

	if client == nil {
		return nil, fmt.Errorf("client nil")
	}

	result := client.HDel(key, filed)
	value, err := result.Result()
	if err != nil {
		logger.Error("err:%v", err)
		return nil, err
	}

	return value, nil
}
