package Kernel

import (
	"errors"
	"github.com/imroc/req"
)

type Client struct {
}

func (this *Client) Get(url string, v ...interface{}) (map[string]interface{}, error) {
	client := req.New()
	response, err := client.Get(url, v)
	if err != nil {
		return nil, err
	}
	data, err := this.parseSuccessResponse(response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this *Client) Post(url string, v ...interface{}) (map[string]interface{}, error) {
	client := req.New()
	response, err := client.Post(url, v)
	if err != nil {
		return nil, err
	}

	return this.parseSuccessResponse(response)
}

func (this *Client) Delete(url string, v ...interface{}) (map[string]interface{}, error) {
	client := req.New()
	response, err := client.Delete(url, v)
	if err != nil {
		return nil, err
	}

	return this.parseSuccessResponse(response)
}

func (this *Client) parseSuccessResponse(response *req.Resp) (map[string]interface{}, error) {
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
