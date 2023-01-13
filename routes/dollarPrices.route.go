package routes

import (
	"dollar-price-server/controllers"

	"github.com/labstack/echo/v4"
)

func DollarPricesRouter(e *echo.Echo) {
	g := e.Group("/dollar-prices")

	// Public routes
	g.GET("", controllers.GetDollarPrices)
	g.POST("", controllers.SaveDollarPrice)

	g.GET("/seed", controllers.SeedDatabase)

}
