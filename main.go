package main

import (
	"dollar-price-server/routes"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {
	runServer()
}

func runServer() {
	loadRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}

func loadRoutes() {
	routes.DollarPricesRouter(e)
}
