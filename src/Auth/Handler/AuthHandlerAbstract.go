package handler

import "context"

type AuthHandlerAbstract struct {
	AuthHandlerInterface
	ctx context.Context
}

var _ AuthHandlerInterface = &AuthHandlerAbstract{}
