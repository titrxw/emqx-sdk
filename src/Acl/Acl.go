package acl

import (
	"context"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
	handler "github.com/titrxw/emqx-sdk/src/Acl/Handler"
)

type Acl struct {
	handler.AclHandlerInterface
	handler handler.AclHandlerInterface
}

func NewAcl(handler handler.AclHandlerInterface) *Acl {
	return &Acl{
		handler: handler,
	}
}

func (this *Acl) Set(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) error {
	return this.handler.Set(ctx, entity, useClientIdType)
}

func (this *Acl) Get(ctx context.Context, clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	return this.handler.Get(ctx, clientName, clientIdType)
}

func (this *Acl) Delete(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) error {
	return this.handler.Delete(ctx, entity, useClientIdType)
}

func (this *Acl) ExportConfig(useClientIdType bool) string {
	return this.handler.ExportConfig(useClientIdType)
}
