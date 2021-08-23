package service

import (
	"context"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"time"
)

type SortService struct {
	logger     *zerolog.Logger
	collection *mongo.Collection
}

func NewSortService(logger *zerolog.Logger, collection *mongo.Collection) *SortService {
	return &SortService{
		logger:     logger,
		collection: collection,
	}
}

func (s SortService) SortData(countries []string, industry string, sector string) (*[]model.ValidTicker, error) {

	filter := bson.M{"locate.country": bson.M{"$in": countries}, "industry": industry, "sector": sector}

	var validTickers []model.ValidTicker

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cur, err := s.collection.Find(ctx, filter)
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

	return &validTickers, err
}
