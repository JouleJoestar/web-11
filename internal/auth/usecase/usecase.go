package usecase

import (
	"web-11/internal/auth/middleware"
	"web-11/internal/auth/provider"
)

type Usecase struct {
	provider *provider.Provider
}

func NewUsecase(prv *provider.Provider) *Usecase {
	return &Usecase{provider: prv}
}

func (uc *Usecase) Register(username, password string) error {
	return uc.provider.CreateUser(username, password)
}

func (uc *Usecase) Login(username, password string) (string, error) {
	existingUser, err := uc.provider.GetUser(username)
	if err != nil {
		return "", err
	}
	return middleware.GenerateJWT(existingUser)
}
