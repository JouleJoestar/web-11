package usecase

type Provider interface {
	GetUser(username string) (string, error)
	CreateUser(username, password string) error
}
