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

func (authHandler *RedisAuthHandler) Set(entity *auth.AuthEntity) {
	var err error
	if entity.GetSalt() == "" {
		err = authHandler.redis.HMSet(authHandler.ctx, authHandler.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword()).Err()
	} else {
		err = authHandler.redis.HMSet(authHandler.ctx, authHandler.clientKeyPrefix+entity.GetClientName(), "password", entity.GetPassword(), "salt", entity.GetSalt()).Err()
	}

	if err != nil {
		panic(err)
	}
}

func (authHandler *RedisAuthHandler) Validate(clientName string, password string) bool {
	sliceCmd := authHandler.redis.HMGet(authHandler.ctx, authHandler.clientKeyPrefix+clientName, "password")
	if sliceCmd.Err() != nil {
		panic(sliceCmd.Err())
	}

	if sliceCmd.Val()[0] == password {
		return true
	}

	return false
}
