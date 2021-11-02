package acl

import (
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

func (this *Acl) Set(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	return this.handler.Set(entity, useClientIdType)
}

func (this *Acl) Get(clientName string, clientIdType string) ([]*entity.AclEntity, error) {
	return this.handler.Get(clientName, clientIdType)
}

func (this *Acl) Delete(entity *entity.AclEntity, useClientIdType bool) (bool, error) {
	return this.handler.Delete(entity, useClientIdType)
}

func (this *Acl) ExportConfig(useClientIdType bool) string {
	return this.handler.ExportConfig(useClientIdType)
}
