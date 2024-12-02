package api

type Usecase interface {
	FetchHelloMessage() (string, error)
	SetHelloMessage(msg string) error
}
