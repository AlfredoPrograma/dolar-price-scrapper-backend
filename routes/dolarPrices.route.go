package routes

import (
	"dolar-price-server/controllers"

	"github.com/labstack/echo/v4"
)

func DolarPricesRoute(e *echo.Echo) {
	g := e.Group("/dollar-prices")

	g.GET("", controllers.GetDolarPrices)
	g.POST("", controllers.SaveDolarPrice)
}
