package usecase

func (u *Usecase) FetchHelloMessage() (string, error) {
	msg, err := u.p.SelectRandomHello()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetHelloMessage(msg string) error {
	isExist, err := u.p.CheckHelloExitByMsg(msg)
	if err != nil {
		return err
	}

	if isExist {
		return nil
	}

	err = u.p.InsertHello(msg)
	if err != nil {
		return err
	}

	return nil
}
