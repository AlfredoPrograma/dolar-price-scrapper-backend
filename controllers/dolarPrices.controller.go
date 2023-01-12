package controllers

import (
	"context"
	"dolar-price-server/common"
	"dolar-price-server/configs"
	"dolar-price-server/models"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dolarPricesCollection = configs.GetCollection(configs.M, "dolar-prices")

func SaveDolarPrice(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var prices common.DolarPrices

	bytesBody, err := io.ReadAll(c.Request().Body)
	json.Unmarshal(bytesBody, &prices)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": "Error reading request body"})
	}

	newDolarPrice := models.DolarPricesModel{
		Id:     primitive.NewObjectID(),
		Prices: prices,
		Date:   time.Now(),
	}

	result, err := dolarPricesCollection.InsertOne(ctx, newDolarPrice)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": "Error inserting data"})
	}

	return c.JSON(http.StatusCreated, &echo.Map{"data": result})
}
