package model

type AuthDomainInterface interface {
	GetEmail() string
	GetPassword() string
}

type Auth struct {
	Email    string
	Password string
}

func NewAuth(email string, password string) *Auth {
	return &Auth{
		Email:    email,
		Password: password,
	}
}

func (a *Auth) GetEmail() string {
	return a.Email
}

func (a *Auth) GetPassword() string {
	return a.Password
}
