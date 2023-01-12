package models

import (
	"dollar-price-server/common"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DollarPrices struct {
	Id     primitive.ObjectID         `json:"id"`
	Prices common.DollarPricesSources `json:"prices"`
	Date   time.Time                  `json:"created_at"`
}
