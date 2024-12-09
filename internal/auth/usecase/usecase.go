package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Usecase struct {
	provider Provider
}

var jwtSecret = []byte("123.456.789")

func NewUsecase(prv Provider) *Usecase {
	return &Usecase{
		provider: prv,
	}
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Токен будет действовать 72 часа
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *Usecase) Register(username, password string) error {
	return uc.provider.CreateUser(username, password)
}

func (uc *Usecase) Login(username, password string) (string, error) {
	username, err := uc.provider.GetUser(username)
	if err != nil {
		return "", err // Возвращаем ошибку, если пользователь не найден
	}

	return GenerateJWT(username) // Генерируем и возвращаем JWT
}
