package handler

import (
	"context"
	"github.com/titrxw/emqx-sdk/src/Acl/Entity"
)

type AclHandlerInterface interface {
	Set(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error)
	Get(ctx context.Context, clientName string, clientIdType string) ([]*entity.AclEntity, error)
	Delete(ctx context.Context, entity *entity.AclEntity, useClientIdType bool) (bool, error)
	ExportConfig(useClientIdType bool) string
}
