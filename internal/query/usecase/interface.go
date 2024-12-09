package usecase

type Provider interface {
	SelectUser(name string) (string, error)
	InsertUser(name string) error
}
