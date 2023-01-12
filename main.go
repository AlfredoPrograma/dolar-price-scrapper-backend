package main

import (
	"dolar-price-server/configs"
	"dolar-price-server/routes"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {
	runServer()
}

func runServer() {
	configs.ConnectDB()

	loadRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}

func loadRoutes() {
	routes.DolarPricesRoute(e)
}
