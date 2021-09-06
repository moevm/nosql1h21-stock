package service

import (
	"context"
	"fmt"
	"nosql1h21-stock-backend/backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	collection *mongo.Collection
}

func NewService(collection *mongo.Collection) *Service {
	return &Service{
		collection: collection,
	}
}

type NoStockInfo struct {
	ticker string
}

func (err NoStockInfo) Error() string {
	return fmt.Sprintf("No info about the stock with ticker %v", err.ticker)
}

func (s *Service) GetStockInfo(ctx context.Context, ticker string) (*model.Stock, error) {
	result := s.collection.FindOne(ctx, bson.M{"symbol": ticker})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, NoStockInfo{ticker: ticker}
		}
		return nil, err
	}
	stock := &model.Stock{}
	if err := result.Decode(stock); err != nil {
		return nil, err
	}
	return stock, nil
}

func (s *Service) findStocks(ctx context.Context, filter interface{}) (stocks []model.StockOverview, _ error) {
	cur, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	stocks = []model.StockOverview{}
	for cur.Next(ctx) {
		var stock model.StockOverview
		err := cur.Decode(&stock)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s *Service) SearchByTicker(ctx context.Context, tickerFragment string) (stocks []model.StockOverview, _ error) {
	return s.findStocks(ctx, bson.M{
		"symbol": bson.M{"$regex": tickerFragment},
	})
}

func (s *Service) SearchByName(ctx context.Context, nameFragment string) (stocks []model.StockOverview, _ error) {
	return s.findStocks(ctx, bson.M{
		"long name": bson.M{"$regex": nameFragment},
	})
}

func (s *Service) getDistinct(ctx context.Context, field string, filter interface{}) ([]string, error) {
	rawValues, err := s.collection.Distinct(ctx, field, filter)
	if err != nil {
		return nil, err
	}

	values := []string{}
	for _, value := range rawValues {
		values = append(values, value.(string))
	}
	return values, nil
}

func (s *Service) GetCountries(ctx context.Context) (countries []string, _ error) {
	return s.getDistinct(ctx, "locate.country", bson.M{})
}

func (s *Service) GetSectors(ctx context.Context) (sectors []string, _ error) {
	return s.getDistinct(ctx, "sector", bson.M{})
}

func (s *Service) GetIndustries(ctx context.Context, sector string) (industries []string, _ error) {
	return s.getDistinct(ctx, "industry", bson.M{"sector": sector})
}

func (s *Service) Filter(ctx context.Context, countries []string, sector, industry string) (stocks []model.StockOverview, _ error) {
	filter := bson.M{}
	if countries != nil {
		filter["locate.country"] = bson.M{"$in": countries}
	}
	if sector != "" {
		filter["sector"] = sector
	}
	if industry != "" {
		filter["industry"] = industry
	}
	return s.findStocks(ctx, filter)
}

type SearchRequest struct {
	Fragment  string // in ticker or company name
	Sector    string
	Industry  string
	Countries []string
}

func (s *Service) Search(ctx context.Context, r SearchRequest) (stocks []model.StockOverview, _ error) {
	filter := bson.M{}
	if r.Countries != nil {
		filter["locate.country"] = bson.M{"$in": r.Countries}
	}
	if r.Sector != "" {
		filter["sector"] = r.Sector
	}
	if r.Industry != "" {
		filter["industry"] = r.Industry
	}
	if r.Fragment != "" {
		filter["$or"] = bson.A{
			bson.M{
				"symbol": bson.M{"$regex": r.Fragment, "$options": "i"},
			},
			bson.M{
				"long name": bson.M{"$regex": r.Fragment, "$options": "i"},
			},
		}
	}
	return s.findStocks(ctx, filter)
}
