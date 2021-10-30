package auth

type AuthEntity struct {
	clientName string
	password   string
	salt       string
}

func (authEntity *AuthEntity) SetClientName(clientName string) {
	authEntity.clientName = clientName
}

func (authEntity *AuthEntity) GetClientName() string {
	return authEntity.clientName
}

func (authEntity *AuthEntity) SetPassword(password string) {
	authEntity.password = password
}

func (authEntity *AuthEntity) GetPassword() string {
	return authEntity.password
}

func (authEntity *AuthEntity) SetSalt(salt string) {
	authEntity.salt = salt
}

func (authEntity *AuthEntity) GetSalt() string {
	return authEntity.salt
}
