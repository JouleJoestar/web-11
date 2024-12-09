package api

type Usecase interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
}
