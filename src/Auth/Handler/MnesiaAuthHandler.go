package handler

import (
	"context"
	"errors"
	"github.com/imroc/req"
	auth "github.com/titrxw/emqx-sdk/src/Auth"
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

func (authHandler *MnesiaAuthHandler) Set(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := authHandler.host + "api/v4/auth_" + authHandler.getAuthClientKeyName(useClientIdType)
	client := req.New()
	header := req.Header{
		"Accept": "application/json",
	}
	param := req.Param{
		authHandler.getAuthClientKeyName(useClientIdType): entity.GetClientName(),
		"password": entity.GetPassword(),
	}
	response, err := client.Post(url, header, param)
	if err != nil {
		return false, err
	}

	_, err = authHandler.parseSuccessResponse(response)
	return true, err
}

func (authHandler *MnesiaAuthHandler) Validate(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := authHandler.host + "api/v4/auth_" + authHandler.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	client := req.New()
	response, err := client.Get(url)
	if err != nil {
		return false, err
	}
	data, err := authHandler.parseSuccessResponse(response)
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

func (authHandler *MnesiaAuthHandler) Delete(entity *auth.AuthEntity, useClientIdType bool) (bool, error) {
	url := authHandler.host + "api/v4/auth_" + authHandler.getAuthClientKeyName(useClientIdType) + "/" + entity.GetClientName()
	client := req.New()
	response, err := client.Delete(url)
	if err != nil {
		return false, err
	}

	_, err = authHandler.parseSuccessResponse(response)
	return true, err
}

func (authHandler *MnesiaAuthHandler) parseSuccessResponse(response *req.Resp) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := response.ToJSON(&data)
	if err != nil {
		return data, err
	}
	_, exists := data["code"]
	if !exists {
		return data, errors.New("emqx 响应失败")
	}
	if data["code"] != 0 {
		return data, errors.New("emqx 响应状态码错误")
	}

	return data, nil
}

func (authHandler *MnesiaAuthHandler) getAuthClientKeyName(useClientIdType bool) string {
	if useClientIdType {
		return "clientid"
	}

	return "username"
}
