package openapi

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
	mqtt "github.com/titrxw/emqx-sdk/src/OpenApi/Mqtt"
)

type OpenApiFactory struct {
	client *kernel.EmqxClient
}

func NewOpenApiFactory(client *kernel.EmqxClient) *OpenApiFactory {
	return &OpenApiFactory{
		client: client,
	}
}

func (this *OpenApiFactory) Message() *mqtt.Message {
	return &mqtt.Message{
		OpenApiAbstract: kernel.OpenApiAbstract{
			Client: this.client,
		},
	}
}
