package controllers

import (
	"context"
	"dollar-price-server/common"
	"dollar-price-server/configs"
	"dollar-price-server/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dollarPricesCollection = configs.GetCollection(configs.M, "dollar-prices")

func SaveDollarPrice(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var prices common.DollarPricesSources

	bytesBody, err := io.ReadAll(c.Request().Body)
	json.Unmarshal(bytesBody, &prices)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": "Error reading request body"})
	}

	newDollarPrice := models.DollarPrices{
		Id:     primitive.NewObjectID(),
		Prices: prices,
		Date:   time.Now(),
	}

	result, err := dollarPricesCollection.InsertOne(ctx, newDollarPrice)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": "Error inserting data"})
	}

	return c.JSON(http.StatusCreated, &echo.Map{"data": result})
}

func GetDollarPrices(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dollarPricesList []models.DollarPrices

	cursor, err := dollarPricesCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": "Error searching data"})
	}

	cursor.All(ctx, &dollarPricesList)

	defer cursor.Close(ctx)
	return c.JSON(200, dollarPricesList)
}
