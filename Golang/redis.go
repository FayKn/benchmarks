package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisPool *redis.Client

func init() {
	redisPool = redis.NewClient(&redis.Options{
		Addr:         getEnvVar("REDIS_HOST") + ":" + getEnvVar("REDIS_PORT"),
		Password:     getEnvVar("REDIS_PASS"),
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 3,
	})
}

func connectToRedis() *redis.Client {
	return redisPool
}

// The rest of your functions remain the same

func getFromRedis(key string) string {
	db := connectToRedis()
	val, err := db.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Error getting value from Redis: ", err, "Key: ", key)
	}
	return val
}

func existsInRedis(key string) bool {
	db := connectToRedis()
	val, err := db.Exists(context.Background(), key).Result()
	if err != nil {
		log.Println("Error checking if key exists in Redis: ", err)
		return false
	}
	return val > 0
}

func setToRedis(key string, value string) {
	db := connectToRedis()
	err := db.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Println("Error setting value in Redis: ", err)
	}
}

func deleteFromRedis(key string) {
	db := connectToRedis()
	err := db.Del(context.Background(), key).Err()
	if err != nil {
		log.Println("Error deleting value from Redis: ", err)
	}
}
