package auth

type AuthEntity struct {
	userName string
	clientId string
	password string
}

func (authEntity *AuthEntity) setUserName(userName string) {
	authEntity.userName = userName
}

func (authEntity *AuthEntity) getUserName() string {
	return authEntity.userName
}

func (authEntity *AuthEntity) setClientId(clientId string) {
	authEntity.clientId = clientId
}

func (authEntity *AuthEntity) getClientId() string {
	return authEntity.clientId
}

func (authEntity *AuthEntity) setPassword(password string) {
	authEntity.password = password
}

func (authEntity *AuthEntity) getPassword() string {
	return authEntity.password
}
