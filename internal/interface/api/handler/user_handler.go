package handler

import (
	"k3_api/internal/interface/api/dto/command"
	"k3_api/internal/interface/api/dto/response"
	"k3_api/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHander struct {
	createUsecase usecase.CreateUser
}

func NewUserHandler(usecase usecase.CreateUser) *UserHander {
	return &UserHander{
		createUsecase: usecase,
	}
}

func (h *UserHander) Create(c echo.Context) error {
	var cmd api.CreateUserCommand
	if err := c.Bind(&cmd); err != nil {
		return err
	}

	user, err := h.createUsecase.Execute(c.Request().Context(), cmd)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(user))
}
