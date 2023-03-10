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
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveDollarPrice(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	prices := common.DollarPricesSources{}

	bytesBody, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{"error": "Error reading request body"})
	}

	json.Unmarshal(bytesBody, &prices)

	newDollarPrice := models.DollarPrices{
		Id:     primitive.NewObjectID(),
		Prices: prices,
		Date:   time.Now(),
	}

	err = newDollarPrice.Save(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": "Error inserting data"})
	}

	return c.JSON(http.StatusCreated, &echo.Map{"data": newDollarPrice})
}

func GetDollarPrices(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dollarPricesList := models.DollarPricesList{}

	err := dollarPricesList.FindAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(200, dollarPricesList)
}

func SeedDatabase(c echo.Context) error {
	if configs.GetEnvVar("CURRENT_ENV") == configs.PROD {
		return c.JSON(http.StatusNotFound, &echo.Map{"ok": false, "msg": "Route not found"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	b, err := os.ReadFile("prices.json")

	if err != nil {
		log.Fatal(err)
	}

	dollarPricesList := models.DollarPricesList{}

	json.Unmarshal(b, &dollarPricesList)

	dollarPricesList.InsertMany(ctx)

	return c.JSON(http.StatusCreated, &echo.Map{"ok": true, "msg": dollarPricesList})
}
