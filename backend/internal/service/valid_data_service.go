package service

import (
	"errors"
	"github.com/rs/zerolog"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
)

type ValidTickersService struct {
	logger *zerolog.Logger
	cache  *repository.ValidTickerCache
}

func NewValidTickersService(logger *zerolog.Logger, cache *repository.ValidTickerCache) *ValidTickersService {
	return &ValidTickersService{
		logger: logger,
		cache:  cache,
	}
}

func (s ValidTickersService) GetValidData() (*model.ValidData, error) {

	if p, ok := s.cache.Load("valid data"); ok {
		s.logger.Info().Msg("Hit cache")
		return &p, nil
	}

	return nil, errors.New("Empty cache for valid data")
}
