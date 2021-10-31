package handler

import (
	"context"
	"github.com/go-redis/redis/v8"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
)

type RedisAuthHandler struct {
	AuthHandlerAbstract
	redis           *redis.Client
	clientKeyPrefix string
}

func NewRedisAuthHandler(ctx context.Context, redis *redis.Client, clientKeyPrefix string) *RedisAuthHandler {
	if clientKeyPrefix == "" {
		clientKeyPrefix = "mqtt:emqx:user:"
	}

	return &RedisAuthHandler{
		redis:           redis,
		clientKeyPrefix: clientKeyPrefix,
		AuthHandlerAbstract: AuthHandlerAbstract{
			ctx: ctx,
		},
	}
}

func (this *RedisAuthHandler) Set(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	var boolCmd *redis.BoolCmd
	if entity.GetSalt() == "" {
		boolCmd = this.redis.HMSet(this.ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword())
	} else {
		boolCmd = this.redis.HMSet(this.ctx, this.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword(), "salt", entity.GetSalt())
	}

	return boolCmd.Val(), boolCmd.Err()
}

func (this *RedisAuthHandler) Validate(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	sliceCmd := this.redis.HMGet(this.ctx, this.clientKeyPrefix+entity.GetClientName(), "password")
	if sliceCmd.Err() != nil {
		return false, sliceCmd.Err()
	}

	return sliceCmd.Val()[0] == entity.GetPassword(), nil
}

func (this *RedisAuthHandler) Delete(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HDel(this.ctx, this.clientKeyPrefix+entity.GetClientName())
	return true, intCmd.Err()
}
