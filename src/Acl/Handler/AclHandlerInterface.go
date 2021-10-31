package handler

import "github.com/titrxw/emqx-sdk/src/Acl"

type AclHandlerInterface interface {
	Set(entity *acl.AclEntity, useClientIdType bool) (bool, error)
	Get(clientName string, clientIdType string) ([]*acl.AclEntity, error)
	Delete(entity *acl.AclEntity, useClientIdType bool) (bool, error)
}
