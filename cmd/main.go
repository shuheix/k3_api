package main

import (
	"k3_api/internal/domain/model"
	"k3_api/internal/infrastructure/persistence"
	"k3_api/internal/interface/api/handler"
	"k3_api/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connnect db")
	}
	db.Debug()
	db.AutoMigrate(&model.User{})

	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	repo := persistence.NewUserRepository(db)
	usecase := usecase.NewCreateUserUsecase(repo)
	handler := handler.NewUserHandler(*usecase)

	routes := e.Group("/api")
	routes.POST("/users", handler.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
