package main

import (
	"context"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
	"sync"
	"time"
)

func GetValidTickers(collection *mongo.Collection, logger *zerolog.Logger, cache *repository.ValidTickerCache, tickersMap *sync.Map) error {
	var validTickers []model.ValidTicker

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Err(err).Send()
		return err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result model.ValidTicker
		err := cur.Decode(&result)
		if err != nil {
			logger.Err(err).Send()
			return err
		}
		validTickers = append(validTickers, result)
	}
	if err := cur.Err(); err != nil {
		logger.Err(err).Send()
		return err
	}

	for _, v := range validTickers {
		tickersMap.Store(v.Symbol, v.ShortName)
	}

	cache.Store("valid tickers", validTickers)
	logger.Info().Msg("Valid tickers store to cache !")

	return nil
}
