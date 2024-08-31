package handlers

import (
	"k3_api/internal/domain/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}


func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

// func (h *UserHandler) GetUsers(c echo.Context) error {
// 	var users []models.User
// 	h.DB.Find(&users)
// 	if h.DB.Error != nil {
// 		return c.String(http.StatusBadRequest, "fail")
// 	}

// 	return c.JSON(http.StatusOK, users)
// }


func (h *UserHandler) GetUsers(c echo.Context) error {
	var usecase.
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var newUser models.CreateUserParams
	err := c.Bind(&newUser); if err != nil {
		return c.JSON(http.StatusBadRequest, newUser)
	}
	user := models.User{}
	user.Name = newUser.Name

	h.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}
