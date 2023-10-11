package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	response, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
	SetKeyValue("set", "ini set", 0)
	value, _ := GetValue("set")
	fmt.Println(value)

	result, _ := DeleteKey("set")
	fmt.Println(result)

	value, _ = GetValue("set")
	fmt.Println(value)

}

// connection
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	DB:       0,
	Password: "",
})

// set key value expired in 5 secound
func SetKeyValue(key string, value interface{}, exparation time.Duration) {
	err := RedisClient.Set(context.Background(), key, value, exparation).Err()
	if err != nil {
		fmt.Println("error set redis")
		return
	}
}

// get value from key
func GetValue(key string) (interface{}, error) {
	value, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil

}

// delete key
func DeleteKey(key string) (int64, error) {
	result, err := RedisClient.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
