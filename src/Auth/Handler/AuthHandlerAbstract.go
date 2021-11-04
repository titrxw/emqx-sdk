package handler

type AuthHandlerAbstract struct {
	AuthHandlerInterface
}

var _ AuthHandlerInterface = &AuthHandlerAbstract{}
