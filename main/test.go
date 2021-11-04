package main

import (
	"github.com/go-redis/redis/v8"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	handler2 "github.com/titrxw/emqx-sdk/src/Acl/Handler"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	handler3 "github.com/titrxw/emqx-sdk/src/Auth/Handler"
)

func main() {
	handler := handler2.NewRedisAclHandler(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379"}), "mqtt:user:")

	authHandler := acl.NewAcl(handler)
	config := authHandler.ExportConfig(false)
	print(config)

	handler1 := handler3.NewRedisAuthHandler(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "mqtt:user:")

	authHandler1 := auth.NewAuth(handler1, nil)
	config1 := authHandler1.ExportConfig(true)
	print(config1)
}
