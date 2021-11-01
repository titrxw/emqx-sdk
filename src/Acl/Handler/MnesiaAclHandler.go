package handler

import (
	"context"
	"errors"
	"github.com/imroc/req"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
	"github.com/titrxw/emqx-sdk/src/Kernel"
	"net/url"
)

type MnesiaAclHandler struct {
	AclHandlerAbstract
	kernel.EmqxClient
}

func NewMnesiaAclHandler(ctx context.Context, host string, appId string, appSecret string) *MnesiaAclHandler {
	return &MnesiaAclHandler{
		AclHandlerAbstract: AclHandlerAbstract{
			ctx: ctx,
		},
		EmqxClient: kernel.EmqxClient{
			Host:      host,
			AppId:     appId,
			AppSecret: appSecret,
		},
	}
}

func (this *MnesiaAclHandler) Set(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	path := "api/v4/acl"

	var err error
	if entity.GetClientName() != "" {
		_, err = this.EmqxClient.Post(path, req.BodyJSON(map[string]string{
			this.getAclClientKeyName(useClientIdType): entity.GetClientName(),
			"topic":  entity.GetTopic(),
			"action": string(entity.GetAction()),
			"access": string(entity.GetAccess()),
		}))
	} else {
		_, err = this.EmqxClient.Post(path, req.BodyJSON(map[string]string{
			"topic":  entity.GetTopic(),
			"action": string(entity.GetAction()),
			"access": string(entity.GetAccess()),
		}))
	}

	if err != nil {
		return false, err
	}
	return true, err
}

func (this *MnesiaAclHandler) Get(clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	var entityMap []*entity.AclEntity
	path := "api/v4/acl/" + clientIdType + "/" + clientName
	data, err := this.EmqxClient.Get(path)
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

func (this *MnesiaAclHandler) Delete(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	var operateType = "$all"
	if entity.GetClientName() != "" {
		operateType = this.getAclClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	}
	path := "api/v4/acl/" + operateType + "/topic/" + url.QueryEscape(entity.GetTopic())
	_, err := this.EmqxClient.Delete(path)
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
