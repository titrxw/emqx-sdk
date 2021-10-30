package handler

type AclHandlerAbstract struct {
	AclHandlerInterface
}

var _ AclHandlerInterface = &AclHandlerAbstract{}
