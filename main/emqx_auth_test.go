package gotest

import (
	"context"
	"github.com/go-redis/redis/v8"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	entity2 "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	handler3 "github.com/titrxw/emqx-sdk/src/Auth/Handler"
	"testing"
)

func TestRedisAuthAdd(t *testing.T) {
	t.Run("testRedisHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(ctx, redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		result, err := authHandler.Set(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestRedisAuthDelete(t *testing.T) {
	t.Run("testRedisHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(ctx, redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		result, err := authHandler.Delete(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthAdd(t *testing.T) {
	t.Run("testMnesiaHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewMnesiaAuthHandler(ctx, "http://127.0.0.1:18083/", "admin", "public")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		entity.SetPassword("test")
		result, err := authHandler.Set(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthDelete(t *testing.T) {
	t.Run("testMnesiaHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewMnesiaAuthHandler(ctx, "http://127.0.0.1:18083/", "admin", "public")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		result, err := authHandler.Delete(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}