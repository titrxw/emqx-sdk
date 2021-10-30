package handler

import (
	"github.com/titrxw/emqx-sdk/src/Auth"
)

type AuthHandlerInterface interface {
	Set(entity *auth.AuthEntity)
	Validate(clientName string, password string) bool
}
