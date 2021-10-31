package auth

import (
	encrypt "github.com/titrxw/emqx-sdk/src/Auth/Encrypt"
	handler "github.com/titrxw/emqx-sdk/src/Auth/Handler"
)

type Auth struct {
	handler.AuthHandlerInterface
	handler handler.AuthHandlerInterface
	encrypt encrypt.EncryptInterface
}

func NewAuth(handler handler.AuthHandlerInterface, encrypt encrypt.EncryptInterface) *Auth {
	return &Auth{
		handler: handler,
		encrypt: encrypt,
	}
}

func (auth *Auth) Set(entity *AuthEntity, useClientIdType bool) (bool, error) {
	if auth.encrypt != nil {
		entity.SetPassword(auth.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return auth.handler.Set(entity, useClientIdType)
}

func (auth *Auth) Validate(entity *AuthEntity, useClientIdType bool) (bool, error) {
	if auth.encrypt != nil {
		entity.SetPassword(auth.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return auth.handler.Validate(entity, useClientIdType)
}

func (auth *Auth) Delete(entity *AuthEntity, useClientIdType bool) (bool, error) {
	return auth.handler.Delete(entity, useClientIdType)
}
