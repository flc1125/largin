package factorys

import "github.com/go-redis/redis/v8"

type connection struct {
	name           string
	Options        *redis.Options
	ClusterOptions *redis.ClusterOptions
}

type connections map[string]connection

func NewRedis(name string) *redis.Client {
	return redis.NewClient(&redis.Options{})
}

// // func NewClusterRedis(name string) *redis.ClusterClient {

// // }
