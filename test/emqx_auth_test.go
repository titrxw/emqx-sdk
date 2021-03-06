package gotest

import (
	"context"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	entity2 "github.com/titrxw/emqx-sdk/src/Auth/Entity"
	handler3 "github.com/titrxw/emqx-sdk/src/Auth/Handler"
	"testing"
)

func TestRedisAuthAdd(t *testing.T) {
	t.Run("testRedisHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(GetRedisClient(), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		err := authHandler.Set(ctx, entity, true)
		if err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestRedisAuthDelete(t *testing.T) {
	t.Run("testRedisHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewRedisAuthHandler(GetRedisClient(), "")

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetPassword("sdfsdfsdf")
		err := authHandler.Delete(ctx, entity, true)
		if err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthAdd(t *testing.T) {
	t.Run("testMnesiaHandlerAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewMnesiaAuthHandler(GetEmqxClient())

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		entity.SetPassword("test")
		err := authHandler.Set(ctx, entity, true)
		if err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAuthDelete(t *testing.T) {
	t.Run("testMnesiaHandlerDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler3.NewMnesiaAuthHandler(GetEmqxClient())

		authHandler := auth.NewAuth(handler, nil)
		entity := new(entity2.AuthEntity)
		entity.SetClientName("test")
		err := authHandler.Delete(ctx, entity, true)
		if err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestExportConfig1(t *testing.T) {
	t.Run("testExportConfig", func(t *testing.T) {
		handler := handler3.NewRedisAuthHandler(nil, "mqtt:user:")

		authHandler := auth.NewAuth(handler, nil)
		config := authHandler.ExportConfig(true)
		if config == "auth.redis.auth_cmd = HMGET mqtt:user:%c password salt" {
			t.Skipped()
		}
		t.Failed()
	})
}
