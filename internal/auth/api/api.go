package api

import (
	"fmt"
	"net/http"
	"time"

	"web-11/internal/auth/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("123.456.789")

type Server struct {
	Address string
	Router  *echo.Echo
	uc      *usecase.Usecase
}

func NewServer(ip string, port int, uc *usecase.Usecase) *Server {
	e := echo.New()
	srv := &Server{
		Address: fmt.Sprintf("%s:%d", ip, port),
		Router:  e,
		uc:      uc,
	}

	srv.Router.POST("/auth/register", srv.Register)
	srv.Router.POST("/auth/login", srv.Login)

	srv.Router.GET("/protected-route", srv.JWTMiddleware(srv.ProtectedRoute))

	return srv
}

// GenerateJWT создает новый JWT-токен для указанного пользователя
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

func (srv *Server) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is required"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		return next(c)
	}
}

func (srv *Server) Register(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Регистрируем пользователя
	err := srv.uc.Register(input.Username, input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Генерируем JWT-токен для нового пользователя
	token, err := GenerateJWT(input.Username) // Генерация токена
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User  registered successfully", "token": token}) // Возвращаем токен
}

func (srv *Server) Login(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	token, err := srv.uc.Login(input.Username, input.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// Пример защищенного маршрута
func (srv *Server) ProtectedRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "This is a protected route!"})
}
