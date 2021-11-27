package handler

import (
	"context"
	"github.com/titrxw/emqx-sdk/src/Auth/Entity"
)

type AuthHandlerInterface interface {
	Set(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error
	Validate(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error
	Delete(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error
	ExportConfig(useClientIdType bool) string
}
