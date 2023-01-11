package main

import (
	"dollar-price-server/configs"
	"dollar-price-server/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func runServer() {
	loadRoutes()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", configs.GetEnvVar("APP_PORT"))))
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
