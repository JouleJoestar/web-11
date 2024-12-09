package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Address   string
	Router    *echo.Echo
	uc        Usecase
	jwtSecret []byte // Добавляем поле для хранения секрета JWT
}

func NewServer(ip string, port int, uc Usecase, jwtSecret string) *Server {
	e := echo.New()
	srv := &Server{
		Address:   fmt.Sprintf("%s:%d", ip, port),
		Router:    e,
		uc:        uc,
		jwtSecret: []byte(jwtSecret), // Инициализируем секрет
	}

	srv.Router.GET("/api/user", srv.JWTMiddleware(srv.GetUser))
	srv.Router.POST("/api/user", srv.JWTMiddleware(srv.PostUser))

	return srv
}

// JWTMiddleware проверяет наличие и валидность JWT-токена
func (srv *Server) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is required"})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return srv.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		return next(c)
	}
}

func (srv *Server) GetUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name parameter is required"})
	}

	user, err := srv.uc.GetUser(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.String(http.StatusOK, "Hello, "+user+"!")
}

func (srv *Server) PostUser(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := srv.uc.CreateUser(input.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Запись добавлена!"})
}
