package api

import (
	"fmt"
	"net/http"
	"web-11/internal/auth/middleware"
	"web-11/internal/auth/usecase"

	"github.com/labstack/echo/v4"
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

	// Определяем маршруты
	srv.Router.POST("/auth/register", srv.Register) // Регистрация (без middleware)
	srv.Router.POST("/auth/login", srv.Login)       // Логин (без middleware)

	// Применяем JWTMiddleware только к защищенным маршрутам
	srv.Router.GET("/protected-route", middleware.JWTMiddleware(srv.ProtectedRoute)) // Пример защищенного маршрута

	return srv
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
	token, err := middleware.GenerateJWT(input.Username) // Генерация токена
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
