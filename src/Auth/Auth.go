package auth

import (
	"context"
	encrypt "github.com/titrxw/emqx-sdk/src/Auth/Encrypt"
	"github.com/titrxw/emqx-sdk/src/Auth/Entity"
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

func (this *Auth) Set(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	if this.encrypt != nil {
		entity.SetPassword(this.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return this.handler.Set(ctx, entity, useClientIdType)
}

func (this *Auth) Validate(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	if this.encrypt != nil {
		entity.SetPassword(this.encrypt.Encode(entity.GetPassword(), entity.GetSalt()))
	}

	return this.handler.Validate(ctx, entity, useClientIdType)
}

func (this *Auth) Delete(ctx context.Context, entity *entity.AuthEntity, useClientIdType bool) error {
	return this.handler.Delete(ctx, entity, useClientIdType)
}

func (this *Auth) ExportConfig(useClientIdType bool) string {
	return this.handler.ExportConfig(useClientIdType)
}
