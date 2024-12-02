package api

import (
	"fmt"
	"web-11/internal/count/usecase"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Address string
	Router  *echo.Echo
	Usecase *usecase.Usecase
}

func NewServer(ip string, port int, use *usecase.Usecase) *Server {
	s := &Server{
		Address: fmt.Sprintf("%s:%d", ip, port),
		Router:  echo.New(),
		Usecase: use,
	}

	s.Router.GET("/count", s.HandleCount)
	s.Router.POST("/count", s.HandleCount)

	return s
}

func (s *Server) HandleCount(c echo.Context) error {
	return s.Usecase.HandleCount(c)
}

func (s *Server) Run() {
	s.Router.Logger.Fatal(s.Router.Start(s.Address))
}
