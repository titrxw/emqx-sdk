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

func (this *RedisAuthHandler) Set(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	var boolCmd *redis.BoolCmd
	if entity.GetSalt() == "" {
		boolCmd = this.redis.HMSet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword())
	} else {
		boolCmd = this.redis.HMSet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword(), "salt", entity.GetSalt())
	}

	return boolCmd.Val(), boolCmd.Err()
}

func (this *RedisAuthHandler) Validate(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	sliceCmd := this.redis.HMGet(ctx, this.clientKeyPrefix+entity.GetClientName(), "password")
	if sliceCmd.Err() != nil {
		return false, sliceCmd.Err()
	}

	return sliceCmd.Val()[0] == entity.GetPassword(), nil
}

func (this *RedisAuthHandler) Delete(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.Del(ctx, this.clientKeyPrefix+entity.GetClientName())
	return true, intCmd.Err()
}

func (this *RedisAuthHandler) ExportConfig(useClientIdType bool) string {
	var typeName = "u"
	if useClientIdType {
		typeName = "c"
	}
	return "auth.redis.auth_cmd = HMGET " + this.clientKeyPrefix + "%" + typeName + " password salt"
}
