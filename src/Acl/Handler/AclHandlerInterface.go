package handler

import "github.com/titrxw/emqx-sdk/src/Acl"

type AclHandlerInterface interface {
	Set(entity acl.AclEntity)
	GetByUserName(userName string) acl.AclEntity
	GetByClientId(clientId string) acl.AclEntity
}
