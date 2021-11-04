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

func NewRedisAclHandler(redis *redis.Client, clientKeyPrefix string) *RedisAclHandler {
	if clientKeyPrefix == "" {
		clientKeyPrefix = "mqtt:emqx:acl:"
	}

	return &RedisAclHandler{
		redis:           redis,
		clientKeyPrefix: clientKeyPrefix,
	}
}

func (this *RedisAclHandler) Set(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HSet(ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic(), string(entity.GetAction()))
	return intCmd.Val() > 0, intCmd.Err()
}

func (this *RedisAclHandler) Get(ctx context.Context, clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	var entityMap []*entity.AclEntity
	maps := this.redis.HGetAll(ctx, this.clientKeyPrefix+clientName)
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

func (this *RedisAclHandler) Delete(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	intCmd := this.redis.HDel(ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic())
	return true, intCmd.Err()
}

func (this *RedisAclHandler) ExportConfig(useClientIdType bool) string {
	var typeName = "u"
	if useClientIdType {
		typeName = "c"
	}

	return "auth.redis.acl_cmd = HGETALL " + this.clientKeyPrefix + "%" + typeName
}
