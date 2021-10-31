package handler

import (
	"github.com/titrxw/emqx-sdk/src/Auth"
)

type AuthHandlerInterface interface {
	Set(entity *auth.AuthEntity, useClientIdType bool) (bool, error)
	Validate(entity *auth.AuthEntity, useClientIdType bool) (bool, error)
	Delete(entity *auth.AuthEntity, useClientIdType bool) (bool, error)
}
