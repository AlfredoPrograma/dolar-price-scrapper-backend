package models

import (
	"dolar-price-scrapper/server/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DolarPricesModel struct {
	Id     primitive.ObjectID `json:"id"`
	Prices common.DolarPrices `json:"prices"`
}
