package service

import (
	"context"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
	"time"
)

type ValidTickersService struct {
	logger     *zerolog.Logger
	cache      *repository.Cache
	collection *mongo.Collection
}

func NewValidTickersService(logger *zerolog.Logger, cache *repository.Cache, collection *mongo.Collection) *ValidTickersService {
	return &ValidTickersService{
		logger:     logger,
		cache:      cache,
		collection: collection,
	}
}

func (s ValidTickersService) GetValidTickers() (*[]model.ValidTicker, error) {

	if p, ok := s.cache.Load("valid tickers"); ok {
		s.logger.Info().Msg("Hit cache")
		return &p, nil
	}

	var validTickers []model.ValidTicker

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

	s.cache.Store("valid tickers", validTickers)
	s.logger.Info().Msg("Store to cache")

	return &validTickers, err
}
