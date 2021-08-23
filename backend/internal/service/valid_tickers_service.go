package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
	"time"
)

type ValidTickersService struct {
	logger     *zerolog.Logger
	cache      *repository.ValidTickerCache
	collection *mongo.Collection
}

func NewValidTickersService(logger *zerolog.Logger, cache *repository.ValidTickerCache, collection *mongo.Collection) *ValidTickersService {
	return &ValidTickersService{
		logger:     logger,
		cache:      cache,
		collection: collection,
	}
}

func (s ValidTickersService) GetValidData() (*model.ValidData, error) {

	if p, ok := s.cache.Load("valid data"); ok {
		s.logger.Info().Msg("Hit cache")
		return &p, nil
	}

	var validData model.ValidData
	var validTickers []model.ValidTicker
	var validSectors []model.Sector

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		s.logger.Err(err).Send()
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result model.ValidTicker
		err := cur.Decode(&result)
		if err != nil {
			s.logger.Err(err).Send()
			break
		}
		validTickers = append(validTickers, result)
	}
	if err := cur.Err(); err != nil {
		s.logger.Err(err).Send()
	}

	sectorsSlice, err := s.collection.Distinct(ctx, "sector", bson.M{})

	if err != nil {
		s.logger.Err(err).Send()
	}

	for _, sector := range sectorsSlice {
		industries, err := s.collection.Distinct(ctx, "industry", bson.M{"sector": sector})
		if err != nil {
			s.logger.Err(err).Send()
			break
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

	validData.Tickers = validTickers
	validData.Sectors = validSectors

	s.cache.Store("valid data", validData)
	s.logger.Info().Msg("Store to cache")

	return &validData, err
}
