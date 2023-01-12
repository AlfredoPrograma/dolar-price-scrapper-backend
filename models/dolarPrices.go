package models

import (
	"dolar-price-server/common"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DolarPricesModel struct {
	Id     primitive.ObjectID `json:"id"`
	Prices common.DolarPrices `json:"prices"`
	Date   time.Time          `json:"created_at"`
}
