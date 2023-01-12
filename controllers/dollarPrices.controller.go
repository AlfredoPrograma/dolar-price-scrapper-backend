package controllers

import (
	"context"
	"dollar-price-server/common"
	"dollar-price-server/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	err = newDollarPrice.Save(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": "Error inserting data"})
	}

	return c.JSON(http.StatusCreated, &echo.Map{"data": newDollarPrice})
}

func GetDollarPrices(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dollarPricesList models.DollarPricesList

	err := dollarPricesList.FindAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(200, dollarPricesList)
}
