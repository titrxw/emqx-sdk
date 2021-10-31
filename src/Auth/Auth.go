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

func (this *Auth) Set(entity *AuthEntity, useClientIdType bool) (bool, error) {
	if this.encrypt != nil {
		entity.SetPassword(this.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return this.handler.Set(entity, useClientIdType)
}

func (this *Auth) Validate(entity *AuthEntity, useClientIdType bool) (bool, error) {
	if this.encrypt != nil {
		entity.SetPassword(this.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return this.handler.Validate(entity, useClientIdType)
}

func (this *Auth) Delete(entity *AuthEntity, useClientIdType bool) (bool, error) {
	return this.handler.Delete(entity, useClientIdType)
}
