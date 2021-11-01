package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
	handler2 "github.com/titrxw/emqx-sdk/src/Acl/Handler"
)

func main() {
	//ctx := context.Background()
	//handler := handler2.NewRedisAuthHandler(ctx, redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//}), "")
	//
	//authHandler := auth.NewAuth(handler, nil)
	//entity := new(entity.AuthEntity)
	//entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
	//entity.SetPassword("sdfsdfsdf")
	//authHandler.Delete(entity, true)

	ctx := context.Background()
	handler := handler2.NewRedisAclHandler(ctx, redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379"}), "")

	authHandler := acl.NewAcl(handler)
	entity := new(entity.AclEntity)
	entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
	entity.SetTopic("/sdf/sdfsdf")
	entity.SetAccessAllow()
	entity.SetActionPub()
	authHandler.Set(entity, true)
	authHandler.Get("lens_z1vX8evgbwuMeb0gbban4GT32ub", "clientid")
}
