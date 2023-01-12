package routes

import (
	"dolar-price-server/controllers"

	"github.com/labstack/echo/v4"
)

func DolarPricesRoute(e *echo.Echo) {
	e.GET("/dolar-prices", controllers.GetDolarPrices)
	e.POST("/dolar-prices", controllers.SaveDolarPrice)
}
