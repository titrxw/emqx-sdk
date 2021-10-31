package handler

import "context"

type AclHandlerAbstract struct {
	AclHandlerInterface
	ctx context.Context
}

var _ AclHandlerInterface = &AclHandlerAbstract{}
