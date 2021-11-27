package handler

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/titrxw/emqx-sdk/src/Auth/Entity"
)

type RedisAuthHandler struct {
	AuthHandlerAbstract
	redis           *redis.Client
	clientKeyPrefix string
}

func NewRedisAuthHandler(redis *redis.Client, clientKeyPrefix string) *RedisAuthHandler {
	if clientKeyPrefix == "" {
		clientKeyPrefix = "mqtt:emqx:user:"
	}

	return &RedisAuthHandler{
		redis:           redis,
		clientKeyPrefix: clientKeyPrefix,
	}
}

func (this *RedisAuthHandler) Set(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	var boolCmd *redis.BoolCmd
	if entity.GetSalt() == "" {
		boolCmd = this.redis.HMSet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword())
	} else {
		boolCmd = this.redis.HMSet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword(), "salt", entity.GetSalt())
	}

	return boolCmd.Err()
}

func (this *RedisAuthHandler) Validate(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	sliceCmd := this.redis.HMGet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password")

	return sliceCmd.Err()
}

func (this *RedisAuthHandler) Delete(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	intCmd := this.redis.Del(ctx, this.clientKeyPrefix+entity.GetClientName())

	return intCmd.Err()
}

func (this *RedisAuthHandler) ExportConfig(useClientIdType bool) string {
	var typeName = "u"
	if useClientIdType {
		typeName = "c"
	}
	return "auth.redis.auth_cmd = HMGET " + this.clientKeyPrefix + "%" + typeName + " password salt"
}
