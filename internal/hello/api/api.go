package api

import (
	"fmt"
	"web-11/internal/auth/middleware"

	"github.com/labstack/echo/v4"
)

type Server struct {
	maxSize int

	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxSize int, uc Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		uc:      uc,
	}

	api.server = echo.New()
	api.server.GET("/hello", middleware.JWTMiddleware(api.GetHello))   // Применяем middleware
	api.server.POST("/hello", middleware.JWTMiddleware(api.PostHello)) // Применяем middleware

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
