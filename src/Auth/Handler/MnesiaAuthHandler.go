package handler

import (
	auth "github.com/titrxw/emqx-sdk/src/Auth"
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

type MnesiaAuthHandler struct {
	AuthHandlerAbstract
	emqxClient kernel.Client
}

func (authHandler *MnesiaAuthHandler) Set(entity auth.AuthEntity) {

}

func (authHandler *MnesiaAuthHandler) Validate(userName string, password string) bool {
	return false
}
