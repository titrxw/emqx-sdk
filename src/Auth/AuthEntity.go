package auth

type AuthEntity struct {
	clientName string
	password   string
	salt       string
}

func (this *AuthEntity) SetClientName(clientName string) {
	this.clientName = clientName
}

func (this *AuthEntity) GetClientName() string {
	return this.clientName
}

func (this *AuthEntity) SetPassword(password string) {
	this.password = password
}

func (this *AuthEntity) GetPassword() string {
	return this.password
}

func (this *AuthEntity) SetSalt(salt string) {
	this.salt = salt
}

func (this *AuthEntity) GetSalt() string {
	return this.salt
}
