package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"web-11/internal/query/usecase"
)

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

	srv.Router.GET("/api/user", srv.GetUser)
	srv.Router.POST("/api/user/create", srv.PostUser)

	return srv
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
