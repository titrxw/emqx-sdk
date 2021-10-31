package handler

import (
	"context"
	"github.com/go-redis/redis/v8"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
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

func (this *RedisAclHandler) Set(entity *acl.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HSet(this.ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic(), entity.GetAction())
	return intCmd.Val() > 0, intCmd.Err()
}

func (this *RedisAclHandler) Get(clientName string, clientIdType string) ([]*acl.AclEntity, error) {
	var entityMap []*acl.AclEntity
	maps := this.redis.HGetAll(this.ctx, this.clientKeyPrefix+clientName)
	if maps.Err() != nil {
		return entityMap, maps.Err()
	}

	for k, v := range maps.Val() {
		entity := new(acl.AclEntity)
		entity.SetClientName(clientName)
		entity.SetTopic(k)
		entity.SetAction(acl.ACL_ACTION(v))
		entity.SetAccessAllow()

		entityMap = append(entityMap, entity)
	}

	return entityMap, nil
}

func (this *RedisAclHandler) Delete(entity *acl.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HDel(this.ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic())
	return true, intCmd.Err()
}
