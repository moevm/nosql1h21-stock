package service

import (
	"context"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
)

type StockService struct {
	logger     *zerolog.Logger
	collection *mongo.Collection
}

func NewStockService(logger *zerolog.Logger, collection *mongo.Collection) *StockService {
	return &StockService{
		logger:     logger,
		collection: collection,
	}
}

func (s StockService) GetAllData(ticker string) (model.Stock, error) {
	filter := bson.D{{"symbol", ticker}}

	result := s.collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return model.Stock{}, result.Err()
	}
	doc := model.Stock{}
	if err := result.Decode(&doc); err != nil {
		s.logger.Err(err).Send()
		return model.Stock{}, err
	}
	return doc, nil
}
