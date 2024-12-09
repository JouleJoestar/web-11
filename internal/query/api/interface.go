package api

type Usecase interface {
	GetUser(name string) (string, error)
	CreateUser(name string) error
}
