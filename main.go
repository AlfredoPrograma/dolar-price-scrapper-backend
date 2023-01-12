package main

import (
	"dollar-price-server/configs"
	"dollar-price-server/routes"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func runServer() {
	loadRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}

func loadRoutes() {
	routes.DollarPricesRouter(e)
}

func main() {
	configs.LoadEnv()

	configs.DB = configs.CreateMongoClient()
	configs.ConnectMongoDB(configs.DB)

	runServer()
}
