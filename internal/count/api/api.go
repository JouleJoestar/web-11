package api

import (
	"fmt"
	"net/http"
	"web-11/internal/auth/middleware"

	"github.com/labstack/echo/v4"
)

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

	s.Router.GET("/count", middleware.JWTMiddleware(s.GetCounter))
	s.Router.POST("/count", middleware.JWTMiddleware(s.UpdateCounter))

	return s
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
