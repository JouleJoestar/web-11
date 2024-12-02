package usecase

import (
	"fmt"
	"net/http"
	"web-11/internal/count/provider"

	"github.com/labstack/echo/v4"
)

type Usecase struct {
	provider *provider.Provider
}

func NewUsecase(prv *provider.Provider) *Usecase {
	return &Usecase{provider: prv}
}

func (u *Usecase) HandleCount(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		counter, err := u.provider.GetCounter()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		err = u.provider.UpdateCounter(1)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.String(http.StatusOK, fmt.Sprintf("%d", counter+1)) // Увеличиваем на 1 для ответа

	case http.MethodPost:
		var requestBody struct {
			Count int `json:"count"`
		}

		if err := c.Bind(&requestBody); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "это не число"})
		}

		err := u.provider.UpdateCounter(requestBody.Count)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})

	default:
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Неизвестный метод"})
	}
}
