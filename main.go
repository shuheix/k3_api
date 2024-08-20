package main

import (
	"github.com/labstack/echo/v4"
	"k3_api/infrastructure/middleware"
	"k3_api/infrastructure/router"
)

func main() {
	e := echo.New()

	middleware.SetUpMiddleware(e)
	router.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
