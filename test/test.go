package gotest

import (
	"github.com/redis/go-redis/v9"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

func GetEmqxClient() *kernel.EmqxClient {
	return &kernel.EmqxClient{
		Host:      "http://127.0.0.1:18083/",
		AppId:     "admin",
		AppSecret: "public",
	}
}

func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379"})
}
