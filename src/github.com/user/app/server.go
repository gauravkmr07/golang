package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/user/app/configs"
	"github.com/user/app/databases"
	"github.com/user/app/routes"
)

func main() {
	fmt.Println("---- Server is preparing to start ----")

	// Setup ENV
	configs.LoadEnv()

	// Connect To Database
	databases.Connect()

	// Start Server
	e := echo.New()

	// Load Routes
	routes.LoadRoutes(e)

	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start("localhost:3001"))
}
