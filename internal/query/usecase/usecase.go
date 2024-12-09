package usecase

import (
	"errors"
)

type Usecase struct {
	defaultMsg string
	p          Provider
}

func NewUsecase(defaultMsg string, p Provider) *Usecase {
	return &Usecase{
		defaultMsg: defaultMsg,
		p:          p,
	}
}

func (u *Usecase) GetUser(name string) (string, error) {
	user, err := u.p.SelectUser(name)
	if err != nil {
		return "", err
	}
	if user == "" {
		return u.defaultMsg, nil
	}
	return user, nil
}

func (u *Usecase) CreateUser(name string) error {
	if name == "" {
		return errors.New("имя пользователя не может быть пустым")
	}
	return nil
}
