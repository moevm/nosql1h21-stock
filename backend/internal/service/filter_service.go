package service

import (
	"context"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nosql1h21-stock-backend/backend/internal/model"
	"time"
)

type FilterService struct {
	logger     *zerolog.Logger
	collection *mongo.Collection
}

func NewFilterService(logger *zerolog.Logger, collection *mongo.Collection) *FilterService {
	return &FilterService{
		logger:     logger,
		collection: collection,
	}
}

func (s FilterService) FilterData(countries []string, sector string, industry string) (*[]model.ValidTicker, error) {

	filter := bson.M{"locate.country": bson.M{"$in": countries}, "sector": sector, "industry": industry}

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
