package handler

import (
	"context"
	"errors"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
	"github.com/titrxw/emqx-sdk/src/Kernel"
	"net/url"
)

type MnesiaAclHandler struct {
	AclHandlerAbstract
	kernel.OpenApiAbstract
}

func NewMnesiaAclHandler(client *kernel.EmqxClient) *MnesiaAclHandler {
	return &MnesiaAclHandler{
		OpenApiAbstract: kernel.OpenApiAbstract{
			Client: client,
		},
	}
}

func (this *MnesiaAclHandler) Set(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	path := "api/v4/acl"

	var err error
	if entity.GetClientName() != "" {
		_, err = this.Client.Post(ctx, path, map[string]string{
			this.getAclClientKeyName(useClientIdType): entity.GetClientName(),
			"topic":  entity.GetTopic(),
			"action": string(entity.GetAction()),
			"access": string(entity.GetAccess()),
		})
	} else {
		_, err = this.Client.Post(ctx, path, map[string]string{
			"topic":  entity.GetTopic(),
			"action": string(entity.GetAction()),
			"access": string(entity.GetAccess()),
		})
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (this *MnesiaAclHandler) Get(ctx context.Context, clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	var entityMap []*entity.AclEntity
	path := "api/v4/acl/" + clientIdType + "/" + clientName
	data, err := this.Client.Get(ctx, path)
	if err != nil {
		return entityMap, err
	}
	_, exists := data["data"]
	if !exists {
		return entityMap, errors.New("emqx 响应数据异常")
	}
	vdata := data["data"].([]interface{})

	for _, v := range vdata {
		content := v.(map[string]interface{})
		aclEntity := new(entity.AclEntity)
		aclEntity.SetClientName(content[clientIdType].(string))
		aclEntity.SetTopic(content["topic"].(string))
		aclEntity.SetAction(entity.ACL_ACTION(content["action"].(string)))
		aclEntity.SetAccess(entity.ACL_ACCESS(content["access"].(string)))
		entityMap = append(entityMap, aclEntity)
	}

	return entityMap, nil
}

func (this *MnesiaAclHandler) Delete(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	var operateType = "$all"
	if entity.GetClientName() != "" {
		operateType = this.getAclClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	}
	path := "api/v4/acl/" + operateType + "/topic/" + url.QueryEscape(entity.GetTopic())
	_, err := this.Client.Delete(ctx, path)
	if err != nil {
		return false, err
	}

	return true, err
}

func (this *MnesiaAclHandler) getAclClientKeyName(useClientIdType bool) string {
	if useClientIdType {
		return "clientid"
	}

	return "username"
}

func (this *MnesiaAclHandler) ExportConfig(useClientIdType bool) string {
	return ""
}
