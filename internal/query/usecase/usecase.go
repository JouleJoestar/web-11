package usecase

import "web-11/internal/query/provider"

type Usecase struct {
	p *provider.Provider
}

func NewUsecase(p *provider.Provider) *Usecase {
	return &Usecase{p: p}
}

func (u *Usecase) GetUser(name string) (string, error) {
	return u.p.SelectUser(name)
}

func (u *Usecase) CreateUser(name string) error {
	return u.p.InsertUser(name)
}
