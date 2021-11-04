package handler

import (
	"context"
	"errors"
	"github.com/imroc/req"
	"github.com/titrxw/emqx-sdk/src/Auth/Entity"
	"github.com/titrxw/emqx-sdk/src/Kernel"
)

type MnesiaAuthHandler struct {
	AuthHandlerAbstract
	kernel.EmqxClient
}

func NewMnesiaAuthHandler(host string, appId string, appSecret string) *MnesiaAuthHandler {
	return &MnesiaAuthHandler{
		EmqxClient: kernel.EmqxClient{
			Host:      host,
			AppId:     appId,
			AppSecret: appSecret,
		},
	}
}

func (this *MnesiaAuthHandler) Set(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	path := "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType)

	_, err := this.EmqxClient.Post(path, req.BodyJSON(map[string]string{
		this.getAuthClientKeyName(useClientIdType): entity.GetClientName(),
		"password": entity.GetPassword(),
	}), ctx)
	if err != nil {
		return false, err
	}
	return true, err
}

func (this *MnesiaAuthHandler) Validate(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	path := "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	data, err := this.EmqxClient.Get(path, ctx)
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

func (this *MnesiaAuthHandler) Delete(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) (bool, error) {
	path := "api/v4/auth_" + this.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	_, err := this.EmqxClient.Delete(path, ctx)
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

func (this *MnesiaAuthHandler) ExportConfig(useClientIdType bool) string {
	return ""
}
