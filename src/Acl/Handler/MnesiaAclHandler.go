package handler

import (
	"context"
	"github.com/imroc/req"
	acl "github.com/titrxw/emqx-sdk/src/Acl"
	"github.com/titrxw/emqx-sdk/src/Kernel"
)

type MnesiaAclHandler struct {
	AclHandlerAbstract
	host string
}

func NewMnesiaAclHandler(ctx context.Context, host string) *MnesiaAclHandler {
	return &MnesiaAclHandler{
		host: host,
		AclHandlerAbstract: AclHandlerAbstract{
			ctx: ctx,
		},
	}
}

func (this *MnesiaAclHandler) Set(entity *acl.AclEntity, useClientIdType bool) (bool, error) {
	url := this.host + "api/v4/acl" + this.getAclClientKeyName(useClientIdType)
	client := new(Kernel.Client)

	var err error
	if entity.GetClientName() != "" {
		_, err = client.Post(url, req.BodyJSON(map[string]string{
			this.getAclClientKeyName(useClientIdType): entity.GetClientName(),
			"topic":  entity.GetTopic(),
			"action": string(entity.GetAction()),
			"access": string(entity.GetAccess()),
		}))
	} else {
		_, err = client.Post(url, req.BodyJSON(map[string]string{
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

func (this *MnesiaAclHandler) Get(clientName string, clientIdType string) ([]*acl.AclEntity, error) {
	var entityMap []*acl.AclEntity
	url := this.host + "api/v4/acl/" + clientIdType + "/" + clientName
	client := new(Kernel.Client)
	data, err := client.Get(url)
	if err != nil {
		return entityMap, err
	}

	for _, v := range data {
		content := v.(map[string]string)
		entity := new(acl.AclEntity)
		entity.SetClientName(content[clientIdType])
		entity.SetTopic(content["topic"])
		entity.SetAction(acl.ACL_ACTION(content["action"]))
		entity.SetAccess(acl.ACL_ACCESS(content["access"]))
		entityMap = append(entityMap, entity)
	}

	return entityMap, nil
}

func (this *MnesiaAclHandler) Delete(entity *acl.AclEntity, useClientIdType bool) (bool, error) {
	var operateType = "$all"
	if entity.GetClientName() != "" {
		operateType = this.getAclClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	}
	url := this.host + "api/v4/acl/" + operateType + "/topic" + entity.GetTopic()
	client := new(Kernel.Client)
	_, err := client.Delete(url)
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
