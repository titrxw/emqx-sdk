package handler

import (
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
)

type AclHandlerInterface interface {
	Set(entity *entity.AclEntity, useClientIdType bool) (bool, error)
	Get(clientName string, clientIdType string) ([]*entity.AclEntity, error)
	Delete(entity *entity.AclEntity, useClientIdType bool) (bool, error)
	ExportConfig(useClientIdType bool) string
}
