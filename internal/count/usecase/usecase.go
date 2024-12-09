package usecase

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type usecase struct {
	provider Provider
}

func NewUsecase(prv Provider) *usecase {
	return &usecase{provider: prv}
}

func (u *usecase) HandleGetCount() (int, error) {
	counter, err := u.provider.GetCounter()
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (u *usecase) HandlePostCount(count int) error {
	return u.provider.UpdateCounter(count)
}

func (u *usecase) HandleGetCountHTTP(c echo.Context) error {
	counter, err := u.HandleGetCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, fmt.Sprintf("%d", counter))
}

func (u *usecase) HandlePostCountHTTP(c echo.Context) error {
	var requestBody struct {
		Count int `json:"count"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "это не число"})
	}

	err := u.HandlePostCount(requestBody.Count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
