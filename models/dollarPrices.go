package models

import (
	"context"
	"dollar-price-server/common"
	"dollar-price-server/configs"
	"dollar-price-server/libs/mongo-utils/collections"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DollarPrices struct {
	Id     primitive.ObjectID         `json:"id"`
	Prices common.DollarPricesSources `json:"prices"`
	Date   time.Time                  `json:"created_at"`
}

type DollarPricesList []DollarPrices

func (m DollarPrices) ToString() string {
	return reflect.TypeOf(m).String()
}

func (l DollarPricesList) ElemToString() string {
	e := reflect.TypeOf(l).Elem()
	return e.String()
}

func (m *DollarPrices) Save(ctx context.Context) error {
	c := collections.GetCollection(configs.DB, collections.StringifyCollectionName(m.ToString()))
	_, err := c.InsertOne(ctx, m)

	if err != nil {
		return err
	}

	return nil
}

func (l *DollarPricesList) FindAll(ctx context.Context) error {
	c := collections.GetCollection(configs.DB, collections.StringifyCollectionName(l.ElemToString()))

	cursor, err := c.Find(ctx, bson.D{})

	if err != nil {
		return err
	}

	err = cursor.All(ctx, l)

	if err != nil {
		return err
	}

	defer cursor.Close(ctx)
	return nil
}

func (l *DollarPricesList) InsertMany(ctx context.Context) error {
	c := collections.GetCollection(configs.DB, collections.StringifyCollectionName(l.ElemToString()))
	v := make([]any, len(*l))

	for i := range v {
		v[i] = (*l)[i]
	}

	_, err := c.InsertMany(ctx, v)

	return err
}
