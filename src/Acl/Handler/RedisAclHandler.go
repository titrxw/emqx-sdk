package handler

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
)

type RedisAclHandler struct {
	AclHandlerAbstract
	redis           *redis.Client
	clientKeyPrefix string
}

func NewRedisAclHandler(ctx context.Context, redis *redis.Client, clientKeyPrefix string) *RedisAclHandler {
	if clientKeyPrefix == "" {
		clientKeyPrefix = "mqtt:emqx:acl:"
	}

	return &RedisAclHandler{
		redis:           redis,
		clientKeyPrefix: clientKeyPrefix,
		AclHandlerAbstract: AclHandlerAbstract{
			ctx: ctx,
		},
	}
}

func (this *RedisAclHandler) Set(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HSet(this.ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic(), string(entity.GetAction()))
	return intCmd.Val() > 0, intCmd.Err()
}

func (this *RedisAclHandler) Get(clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	var entityMap []*entity.AclEntity
	maps := this.redis.HGetAll(this.ctx, this.clientKeyPrefix+clientName)
	if maps.Err() != nil {
		return entityMap, maps.Err()
	}

	for k, v := range maps.Val() {
		aclEntity := new(entity.AclEntity)
		aclEntity.SetClientName(clientName)
		aclEntity.SetTopic(k)
		aclEntity.SetAction(entity.ACL_ACTION(v))
		aclEntity.SetAccessAllow()

		entityMap = append(entityMap, aclEntity)
	}

	return entityMap, nil
}

func (this *RedisAclHandler) Delete(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HDel(this.ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic())
	return true, intCmd.Err()
}
