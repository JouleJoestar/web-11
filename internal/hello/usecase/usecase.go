package usecase

type Usecase struct {
	defaultMsg string

	p Provider
}

func NewUsecase(defaultMsg string, p Provider) *Usecase {
	return &Usecase{
		defaultMsg: defaultMsg,
		p:          p,
	}
}
