package main

import (
	"dolar-price-scrapper/server"
	"dolar-price-scrapper/server/configs"
)

func main() {
	// scrapper.GetDolarPrice()

	configs.ConnectDB()
	server.RunServer()
}
