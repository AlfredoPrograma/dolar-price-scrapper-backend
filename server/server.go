package server

import (
	"dolar-price-scrapper/server/routes"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func RunServer() {
	loadRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}

func loadRoutes() {
	routes.DolarPricesRoute(e)
}
