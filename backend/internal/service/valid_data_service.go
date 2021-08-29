package service

import (
	"errors"
	"github.com/rs/zerolog"
	"nosql1h21-stock-backend/backend/internal/model"
	"nosql1h21-stock-backend/backend/internal/repository"
)

type ValidDataService struct {
	logger *zerolog.Logger
	cache  *repository.ValidTickerCache
}

func NewValidDataService(logger *zerolog.Logger, cache *repository.ValidTickerCache) *ValidDataService {
	return &ValidDataService{
		logger: logger,
		cache:  cache,
	}
}

func (s ValidDataService) GetValidData() (*model.ValidData, error) {

	if p, ok := s.cache.Load("valid data"); ok {
		s.logger.Info().Msg("Hit cache")
		return &p, nil
	}

	return nil, errors.New("Empty cache for valid data")
}
