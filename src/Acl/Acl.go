package acl

import (
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

func (this *Acl) Set(entity *AclEntity, useClientIdType bool) (bool, error) {
	return this.handler.Set(entity, useClientIdType)
}

func (this *Acl) Get(clientName string, clientIdType string) ([]*AclEntity, error) {
	return this.handler.Get(clientName, clientIdType)
}

func (this *Acl) Delete(entity *AclEntity, useClientIdType bool) (bool, error) {
	return this.handler.Delete(entity, useClientIdType)
}
