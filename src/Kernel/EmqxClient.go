package kernel

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/imroc/req"
)

type EmqxClient struct {
	Host      string
	AppId     string
	AppSecret string
}

func NewClient(host string, appId string, appSecret string) *EmqxClient {
	return &EmqxClient{
		Host:      host,
		AppId:     appId,
		AppSecret: appSecret,
	}
}

func (this *EmqxClient) Get(ctx context.Context, path string) (map[string]interface{}, error) {
	response, err := req.Get(this.Host+path, ctx, this.getAuthorizationHeader())
	if err != nil {
		return nil, err
	}
	data, err := this.parseSuccessResponse(response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this *EmqxClient) Post(ctx context.Context, path string, params interface{}) (map[string]interface{}, error) {
	response, err := req.Post(this.Host+path, ctx, this.getAuthorizationHeader(), req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	return this.parseSuccessResponse(response)
}

func (this *EmqxClient) Delete(ctx context.Context, path string) (map[string]interface{}, error) {
	response, err := req.Delete(this.Host+path, ctx, this.getAuthorizationHeader())
	if err != nil {
		return nil, err
	}

	return this.parseSuccessResponse(response)
}

func (this *EmqxClient) getAuthorizationHeader() req.Header {
	auth := this.AppId + ":" + this.AppSecret

	return req.Header{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(auth)),
	}
}

func (this *EmqxClient) parseSuccessResponse(response *req.Resp) (map[string]interface{}, error) {
	if response.Response().StatusCode != 200 {
		msg, _ := response.ToString()
		return nil, errors.New(msg)
	}

	var data map[string]interface{}
	err := response.ToJSON(&data)
	if err != nil {
		return data, err
	}
	_, exists := data["code"]
	if !exists {
		return data, errors.New("emqx 响应失败")
	}
	if data["code"].(float64) != 0 {
		return data, errors.New("emqx 响应状态码错误")
	}

	return data, nil
}
