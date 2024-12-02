package usecase

type Provider interface {
	SelectRandomHello() (string, error)
	CheckHelloExitByMsg(string) (bool, error)
	InsertHello(string) error
}
