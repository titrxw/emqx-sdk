package mqtt

import (
	"context"
	"github.com/imroc/req"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

type Message struct {
	kernel.OpenApiAbstract
}

func (this *Message) publish(ctx context.Context, topic string, clientId string, payload string, qos int, retain bool) (bool, error) {
	_, err := this.Client.Post("api/v4/mqtt/publish", req.BodyJSON(map[string]interface{}{
		"topic":    topic,
		"clientid": clientId,
		"payload":  payload,
		"qos":      qos,
		"retain":   retain,
	}), ctx)
	if err != nil {
		return false, err
	}
	return true, err
}
