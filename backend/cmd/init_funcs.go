package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
	"sync"
	"time"
)

func GetValidData(collection *mongo.Collection, logger *zerolog.Logger, cache *repository.ValidTickerCache, tickersMap *sync.Map) error {
	var validData model.ValidData
	var validTickers []model.ValidTicker
	var validSectors []model.Sector

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

	sectorsSlice, err := collection.Distinct(ctx, "sector", bson.M{})

	if err != nil {
		logger.Err(err).Send()
		return err
	}

	for _, sector := range sectorsSlice {
		industries, err := collection.Distinct(ctx, "industry", bson.M{"sector": sector})
		if err != nil {
			logger.Err(err).Send()
			return err
		}

		validIndustries := make([]string, len(industries))
		for i, v := range industries {
			validIndustries[i] = fmt.Sprint(v)
		}

		validSectors = append(validSectors, model.Sector{
			Sector:     sector.(string),
			Industries: validIndustries,
		})
	}

	if err == nil {
		validData.Tickers = validTickers
		validData.Sectors = validSectors

		for _, v := range validTickers {
			tickersMap.Store(v.Symbol, v.ShortName)
		}

		cache.Store("valid data", validData)
		logger.Info().Msg("Store to cache valid data")
	}

	return err
}
