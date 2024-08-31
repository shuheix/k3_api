package main

import (
	"k3_api/internal/domain/models"
	"k3_api/internal/domain/repository"
	"k3_api/internal/infrastructure/repository"
	"k3_api/internal/presentation/handlers"
	usecase "k3_api/internal/usecase/user"

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

	e.Logger.Fatal(e.Start(":1323"))
}