package gotest

import (
	"context"
	"github.com/go-redis/redis/v8"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	entity "github.com/titrxw/emqx-sdk/src/Acl/Entity"
	handler2 "github.com/titrxw/emqx-sdk/src/Acl/Handler"
	"testing"
)

func TestRedisAclAdd(t *testing.T) {
	t.Run("testRedisAclAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler2.NewRedisAclHandler(ctx, redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379"}), "")

		aclHandler := acl.NewAcl(handler)
		entity := new(entity.AclEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetTopic("/sdf/sdfsdf")
		entity.SetAccessAllow()
		entity.SetActionPub()
		result, err := aclHandler.Set(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestRedisAclDelete(t *testing.T) {
	t.Run("testRedisAclDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler2.NewRedisAclHandler(ctx, redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379"}), "")

		aclHandler := acl.NewAcl(handler)
		entity := new(entity.AclEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetTopic("/sdf/sdfsdf")
		entity.SetAccessAllow()
		entity.SetActionPub()
		result, err := aclHandler.Delete(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAclAdd(t *testing.T) {
	t.Run("testMnesiaAclAdd", func(t *testing.T) {
		ctx := context.Background()
		handler := handler2.NewMnesiaAclHandler(ctx, "http://127.0.0.1:18083/", "admin", "public")

		aclHandler := acl.NewAcl(handler)
		entity := new(entity.AclEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetTopic("/sdf/sdfsdf")
		entity.SetAccessAllow()
		entity.SetActionPub()
		result, err := aclHandler.Set(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestMnesiaAclDelete(t *testing.T) {
	t.Run("testMnesiaAclDelete", func(t *testing.T) {
		ctx := context.Background()
		handler := handler2.NewMnesiaAclHandler(ctx, "http://127.0.0.1:18083/", "admin", "public")

		aclHandler := acl.NewAcl(handler)
		entity := new(entity.AclEntity)
		entity.SetClientName("lens_z1vX8evgbwuMeb0gbban4GT32ub")
		entity.SetTopic("/sdf/sdfsdf")
		entity.SetAccessAllow()
		entity.SetActionPub()
		result, err := aclHandler.Delete(entity, true)
		if !result || err != nil {
			t.Failed()
		}
		t.Skipped()
	})
}

func TestExportConfig(t *testing.T) {
	t.Run("testExportConfig", func(t *testing.T) {
		ctx := context.Background()
		handler := handler2.NewRedisAclHandler(ctx, redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379"}), "mqtt:user:")

		aclHandler := acl.NewAcl(handler)
		config := aclHandler.ExportConfig(true)
		if config == "auth.redis.acl_cmd = HGETALL mqtt:user:%c" {
			t.Skipped()
		}
		t.Failed()
	})
}