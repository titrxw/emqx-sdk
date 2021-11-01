package handler

import (
	"github.com/titrxw/emqx-sdk/src/Auth/Entity"
)

type AuthHandlerInterface interface {
	Set(entity *entity.AuthEntity, useClientIdType bool) (bool, error)
	Validate(entity *entity.AuthEntity, useClientIdType bool) (bool, error)
	Delete(entity *entity.AuthEntity, useClientIdType bool) (bool, error)
}
