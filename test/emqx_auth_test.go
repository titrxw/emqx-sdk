package gotest

import (
	"context"
	"github.com/go-redis/redis/v8"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	entity2 "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	handler3 "github.com/titrxw/emqx-sdk/src/Auth/Handler"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	"testing"
)

func TestRedisAuthAdd(t *testing.T) {
	t.Run("testRedisHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		result, err := authHandler.Set(ctx, entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestRedisAuthDelete(t *testing.T) {
	t.Run("testRedisHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		result, err := authHandler.Delete(ctx, entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthAdd(t *testing.T) {
	t.Run("testMnesiaHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		client := &kernel.EmqxClient{
			Host:      "http://127.0.0.1:18083/",
			AppId:     "admin",
			AppSecret: "public",
		}
		handler := handler3.NewMnesiaAuthHandler(client)

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		entity.SetPassword("test")
		result, err := authHandler.Set(ctx, entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthDelete(t *testing.T) {
	t.Run("testMnesiaHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		client := &kernel.EmqxClient{
			Host:      "http://127.0.0.1:18083/",
			AppId:     "admin",
			AppSecret: "public",
		}
		handler := handler3.NewMnesiaAuthHandler(client)

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		result, err := authHandler.Delete(ctx, entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestExportConfig1(t *testing.T) {
	t.Run("testExportConfig", func(t *testing.T) {
		handler := handler3.NewRedisAuthHandler(redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}), "mqtt:user:")

		authHandler := auth.NewAuth(handler, nil)
		config := authHandler.ExportConfig(true)
		if config == "auth.redis.auth_cmd = HMGET mqtt:user:%c password salt" {
			t.Skipped()
		}
		t.Failed()
	})
}
