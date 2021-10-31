package handler

import (
	"context"
	"errors"
	"github.com/imroc/req"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	"github.com/titrxw/emqx-sdk/src/Kernel"
)

type MnesiaAuthHandler struct {
	AuthHandlerAbstract
	host      string
	appId     string
	appSecret string
}

func NewMnesiaAuthHandler(ctx context.Context, host string, appId string, appSecret string) *MnesiaAuthHandler {
	return &MnesiaAuthHandler{
		host:      host,
		appId:     appId,
		appSecret: appSecret,
		AuthHandlerAbstract: AuthHandlerAbstract{
			ctx: ctx,
		},
	}
}

func (this *MnesiaAuthHandler) Set(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := this.host + "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType)
	client := new(Kernel.Client)

	_, err := client.Post(url, req.BodyJSON(map[string]string{
		this.getAuthClientKeyName(useClientIdType): entity.GetClientName(),
		"password": entity.GetPassword(),
	}))
	if err != nil {
		return false, err
	}
	return true, err
}

func (this *MnesiaAuthHandler) Validate(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := this.host + "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	client := new(Kernel.Client)
	data, err := client.Get(url)
	if err != nil {
		return false, err
	}

	_, exists := data["data"]
	if !exists {
		return false, errors.New("emqx 响应数据异常")
	}
	content := data["data"].(map[string]string)
	_, exists = content["password"]
	if !exists {
		return false, errors.New("emqx 响应数据异常")
	}
	if content["password"] != entity.GetPassword() {
		return false, errors.New("密码错误")
	}

	return true, nil
}

func (this *MnesiaAuthHandler) Delete(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := this.host + "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	client := new(Kernel.Client)
	_, err := client.Delete(url)
	if err != nil {
		return false, err
	}

	return true, err
}

func (this *MnesiaAuthHandler) getAuthClientKeyName(useClientIdType bool) string {
	if useClientIdType {
		return "clientid"
	}

	return "username"
}
