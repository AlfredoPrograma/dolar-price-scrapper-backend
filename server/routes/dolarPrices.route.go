package routes

import (
	"dolar-price-scrapper/server/controllers"

	"github.com/labstack/echo/v4"
)

func DolarPricesRoute(e *echo.Echo) {
	e.POST("/dolar-prices", controllers.SaveDolarPrice)
}
