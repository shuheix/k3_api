package main

import (
	"k3_api/internal/domain/models"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userService struct {
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connnect db")
	}
	db.Debug()
	db.AutoMigrate(&models.User{})
	

	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/users", func(c echo.Context) error {
		var users []models.User
		db.Find(&users)
		if db.Error != nil {
			return c.String(http.StatusBadRequest, "fail")
		}

		return c.JSON(http.StatusOK, users)
	})
	e.POST("/users", func(c echo.Context) error {
		var newUser models.CreateUserParams
		err := c.Bind(&newUser)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad")
		}

		var user models.User
		user.Name = newUser.Name

		spew.Dump(user)

		result := db.Create(&user)
		if result.Error != nil {
			return c.String(http.StatusConflict, "bad")
		}
		return c.JSON(http.StatusOK, user)

	})

	e.Logger.Fatal(e.Start(":1323"))
}
