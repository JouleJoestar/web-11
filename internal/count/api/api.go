package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("123.456.789")

type Server struct {
	Address string
	Router  *echo.Echo
	Usecase Usecase
}

func NewServer(ip string, port int, usecase Usecase) *Server {
	s := &Server{
		Address: fmt.Sprintf("%s:%d", ip, port),
		Router:  echo.New(),
		Usecase: usecase,
	}

	s.Router.GET("/count", s.JWTMiddleware(s.GetCounter))
	s.Router.POST("/count", s.JWTMiddleware(s.UpdateCounter))

	return s
}

func (s *Server) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

func (s *Server) GetCounter(c echo.Context) error {
	count, err := s.Usecase.HandleGetCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, fmt.Sprintf("%d", count))
}

func (s *Server) UpdateCounter(c echo.Context) error {
	var requestBody struct {
		Count int `json:"count"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "это не число"})
	}

	err := s.Usecase.HandlePostCount(requestBody.Count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}

func (s *Server) Run() {
	s.Router.Logger.Fatal(s.Router.Start(s.Address))
}
