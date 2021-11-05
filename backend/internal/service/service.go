package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"nosql1h21-stock-backend/backend/internal/model"
	"strconv"
	"strings"

	. "github.com/gobeam/mongo-go-pagination"
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

func (s *Service) Export(ctx context.Context) (jsonEncoded []byte, _ error) {
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, nil
	}
	defer cur.Close(ctx)

	stocks := []model.Stock{}
	for cur.Next(ctx) {
		var stock model.Stock
		err := cur.Decode(&stock)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return json.Marshal(stocks)
}

func (s *Service) Import(ctx context.Context, jsonEncoded io.Reader) error {
	err := s.collection.Drop(ctx)
	if err != nil {
		return err
	}

	var stocks []model.Stock
	err = json.NewDecoder(jsonEncoded).Decode(&stocks)
	if err != nil {
		return err
	}

	var documents []interface{}
	for _, stock := range stocks {
		documents = append(documents, stock)
	}

	_, err = s.collection.InsertMany(ctx, documents)
	return err
}

type CountItem struct {
	Key    string `bson:"_id"`
	Amount float64
	Unit   string
}

type ErrInvalidArgument struct {
	Arg string
}

func (e ErrInvalidArgument) Error() string {
	if e.Arg != "" {
		e.Arg = " '" + e.Arg + "'"
	}
	return "Invalid argument" + e.Arg
}

func (s *Service) AggregateCountCompanies(ctx context.Context, in string, filter FilterRequest) ([]CountItem, error) {
	mongoFilter := bson.M{}
	setFilter(filter, mongoFilter)

	switch in {
	case "sector", "industry":
		break
	case "country":
		in = "locate.country"
	default:
		return nil, ErrInvalidArgument{"in"}
	}

	cur, err := s.collection.Aggregate(ctx, bson.A{
		bson.M{"$match": mongoFilter},
		bson.M{"$group": bson.M{"_id": "$" + in, "amount": bson.M{"$sum": 1}}},
		bson.M{"$sort": bson.M{"amount": -1}},
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	items := []CountItem{}
	for cur.Next(ctx) {
		var item CountItem
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) AggregateAverage(ctx context.Context, property, in string, filter FilterRequest) ([]CountItem, error) {
	mongoFilter := bson.M{}
	setFilter(filter, mongoFilter)

	switch in {
	case "sector", "industry":
		break
	case "country":
		in = "locate.country"
	default:
		return nil, ErrInvalidArgument{"in"}
	}

	switch property {
	case "employees":
		property = "staff.employees"
	case "quick ratio":
		property = "financial data.quick ratio"
	case "current ratio":
		property = "financial data.current ratio"
	case "debt to equity":
		property = "financial data.debt to equity"
	case "roa":
		property = "financial data.roa"
	case "roe":
		property = "financial data.roe"
	default:
		return nil, ErrInvalidArgument{"property"}
	}

	cur, err := s.collection.Aggregate(ctx, bson.A{
		bson.M{"$match": mongoFilter},
		bson.M{"$group": bson.M{"_id": "$" + in, "amount": bson.M{"$avg": "$" + property}}},
		bson.M{"$sort": bson.M{"amount": -1}},
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	items := []CountItem{}
	for cur.Next(ctx) {
		var item CountItem
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Aggregate(ctx context.Context, mode, property, in string, filter FilterRequest) ([]CountItem, error) {
	switch mode {
	case "count":
		return s.AggregateCountCompanies(ctx, in, filter)
	case "average":
		return s.AggregateAverage(ctx, property, in, filter)
	}
	return nil, ErrInvalidArgument{"mode"}
}

type FilterRequest struct {
	SectorFilter      string
	IndustryFilter    string
	EmployeesFilter   string
	CountriesFilter   []string
	TotalCash         string
	TotalCashPerShare string
	Ebitda            string
	TotalDebt         string
	QuickRatio        string
	CurrentRatio      string
	TotalRevenue      string
	RevenuePerShare   string
	DebtToEquity      string
	ReturnOnAssets    string
	ReturnOnEquity    string
}

func (s *Service) filerStocks(ctx context.Context, filter interface{}, page int64) (model.TableData, error) {

	projection := bson.D{
		{"symbol", 1},
		{"short name", 1},
		{"industry", 1},
		{"sector", 1},
		{"staff.employees", 1},
		{"locate.country", 1},
		{"financial data.total cash", 1},
		{"financial data.total cash per share", 1},
		{"financial data.ebitda", 1},
		{"financial data.total debt", 1},
		{"financial data.quick ratio", 1},
		{"financial data.current ratio", 1},
		{"financial data.total revenue", 1},
		{"financial data.revenue per share", 1},
		{"financial data.debt to equity", 1},
		{"financial data.roa", 1},
		{"financial data.roe", 1},
		{"financial data.financial currency", 1},
	}

	var limit int64 = 40

	stocks := []model.TableFilterData{}
	paginatedData, err := New(s.collection).Context(ctx).Limit(limit).Page(page).Sort("symbol", 1).Select(projection).Filter(filter).Decode(&stocks).Find()

	if err != nil {
		return model.TableData{}, err
	}

	return model.TableData{
		Stocks:    stocks,
		Page:      page,
		TotalPage: paginatedData.Pagination.TotalPage,
	}, nil
}

func setFilterForIntValue(value string, filterValue string, filter *bson.M) {
	if value[0:1] == ">" {
		employeesCountString := strings.ReplaceAll(value, ">", "")
		employeesCount, err := strconv.Atoi(employeesCountString)
		if err == nil {
			filter := *(filter)
			filter[filterValue] = bson.M{"$gte": employeesCount}
		}
	} else if value[0:1] == "<" {
		employeesCountString := strings.ReplaceAll(value, "<", "")
		employeesCount, err := strconv.Atoi(employeesCountString)
		if err == nil {
			filter := *(filter)
			filter[filterValue] = bson.M{"$lte": employeesCount}
		}
	} else {
		employeesCount, err := strconv.Atoi(value)
		if err == nil {
			filter := *(filter)
			filter[filterValue] = bson.M{"$eq": employeesCount}
		}
	}
}

func setFilterForFloatValue(value string, filterValue string, filter *bson.M) {
	if value[0:1] == ">" {
		employeesCountString := strings.ReplaceAll(value, ">", "")
		employeesCount, err := strconv.ParseFloat(employeesCountString, 32)
		if err == nil {
			filter := *(filter)
			filter[filterValue] = bson.M{"$gte": employeesCount}
		}
	} else if value[0:1] == "<" {
		employeesCountString := strings.ReplaceAll(value, "<", "")
		employeesCount, err := strconv.ParseFloat(employeesCountString, 32)
		if err == nil {
			filter := *(filter)
			filter[filterValue] = bson.M{"$lte": employeesCount}
		}
	} else {
		employeesCount, err := strconv.ParseFloat(value, len(value))
		if err == nil {
			filter := *(filter)
			filter[filterValue] = employeesCount
		}
	}
}

func setFilter(r FilterRequest, filter bson.M) {
	if r.SectorFilter != "" {
		filter["sector"] = r.SectorFilter
	}

	if r.IndustryFilter != "" {
		filter["industry"] = r.IndustryFilter
	}

	if r.CountriesFilter != nil {
		filter["locate.country"] = bson.M{"$in": r.CountriesFilter}
	}

	if r.EmployeesFilter != "" {
		setFilterForIntValue(r.EmployeesFilter, "staff.employees", &filter)
	}

	if r.TotalCash != "" {
		setFilterForIntValue(r.TotalCash, "financial data.total cash", &filter)
	}

	if r.TotalCashPerShare != "" {
		setFilterForFloatValue(r.TotalCashPerShare, "financial data.total cash per share", &filter)
	}

	if r.Ebitda != "" {
		setFilterForIntValue(r.Ebitda, "financial data.ebitda", &filter)
	}

	if r.TotalDebt != "" {
		setFilterForIntValue(r.TotalDebt, "financial data.total debt", &filter)
	}

	if r.QuickRatio != "" {
		setFilterForFloatValue(r.QuickRatio, "financial data.quick ratio", &filter)
	}

	if r.CurrentRatio != "" {
		setFilterForFloatValue(r.CurrentRatio, "financial data.current ratio", &filter)
	}

	if r.TotalRevenue != "" {
		setFilterForIntValue(r.TotalRevenue, "financial data.total revenue", &filter)
	}

	if r.RevenuePerShare != "" {
		setFilterForFloatValue(r.RevenuePerShare, "financial data.revenue per share", &filter)
	}

	if r.DebtToEquity != "" {
		setFilterForFloatValue(r.DebtToEquity, "financial data.debt to equity", &filter)
	}

	if r.ReturnOnAssets != "" {
		setFilterForFloatValue(r.ReturnOnAssets, "financial data.roa", &filter)
	}

	if r.ReturnOnEquity != "" {
		setFilterForFloatValue(r.ReturnOnEquity, "financial data.roe", &filter)
	}
}

func (s *Service) TableFilter(ctx context.Context, r FilterRequest, page int64) (model.TableData, error) {
	filter := bson.M{}

	setFilter(r, filter)

	return s.filerStocks(ctx, filter, page)
}
