package handler

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
)

type RedisAclHandler struct {
	AclHandlerAbstract
	redis           redis.Cmdable
	clientKeyPrefix string
}

func NewRedisAclHandler(redis redis.Cmdable, clientKeyPrefix string) *RedisAclHandler {
	if clientKeyPrefix == "" {
		clientKeyPrefix = "mqtt:emqx:acl:"
	}

	return &RedisAclHandler{
		redis:           redis,
		clientKeyPrefix: clientKeyPrefix,
	}
}

func (this *RedisAclHandler) Set(ctx context.Context, aclEntity *entity.AclEntity, useClientIdType bool) error {
	permission := ""
	switch aclEntity.GetAction() {
	case entity.ACTION_SUB:
		permission = "1"
	case entity.ACTION_PUB:
		permission = "2"
	case entity.ACTION_PUBSUB:
		permission = "3"
	}
	intCmd := this.redis.HSet(ctx, this.clientKeyPrefix+aclEntity.GetClientName(), aclEntity.GetTopic(), permission)

	return intCmd.Err()
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
		action := entity.ACL_ACTION("")
		switch v {
		case "1":
			action = entity.ACTION_SUB
		case "2":
			action = entity.ACTION_PUB
		case "3":
			action = entity.ACTION_PUBSUB

		}
		aclEntity.SetAction(action)
		aclEntity.SetAccessAllow()

		entityMap = append(entityMap, aclEntity)
	}

	return entityMap, nil
}

func (this *RedisAclHandler) Delete(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) error {
	intCmd := this.redis.HDel(ctx, this.clientKeyPrefix+entity.GetClientName(), entity.GetTopic())

	return intCmd.Err()
}

func (this *RedisAclHandler) ExportConfig(useClientIdType bool) string {
	var typeName = "u"
	if useClientIdType {
		typeName = "c"
	}

	return "auth.redis.acl_cmd = HGETALL " + this.clientKeyPrefix + "%" + typeName
}
