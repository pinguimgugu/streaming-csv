package main

import (
	"github.com/pinguimgugu/streaming-csv/application/routes"
	"github.com/pinguimgugu/streaming-csv/infrastructure/factory"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	routes.NewUserRoutes(e, factory.GetUserRepository()).Handler()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
