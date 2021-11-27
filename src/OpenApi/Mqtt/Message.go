package mqtt

import (
	"context"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

type Message struct {
	kernel.OpenApiAbstract
}

func (this *Message) Publish(ctx context.Context, topic string, clientId string, payload string, qos int, retain bool) error {
	_, err := this.Client.Post(ctx, "api/v4/mqtt/publish", map[string]interface{}{
		"topic":    topic,
		"clientid": clientId,
		"payload":  payload,
		"qos":      qos,
		"retain":   retain,
	})

	return err
}
