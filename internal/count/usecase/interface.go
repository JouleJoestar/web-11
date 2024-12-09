package usecase

type Provider interface {
	GetCounter() (int, error)
	UpdateCounter(count int) error
}
