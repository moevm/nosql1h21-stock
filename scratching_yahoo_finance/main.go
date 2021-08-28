package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
	"time"

	"nosql1h21-stock/scratching_yahoo_finance/requests"
	"nosql1h21-stock/scratching_yahoo_finance/tickers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Stock struct {
	Symbol        string `bson:"symbol"`
	ShortName     string `bson:"short name,omitempty"`
	LongName      string `bson:"long name,omitempty"`
	Summary       string `bson:"summary,omitempty"`
	Industry      string `bson:"industry,omitempty"`
	Sector        string `bson:"sector,omitempty"`
	Staff         Staff
	Locate        Locate
	Contacts      Contacts
	FinancialData FinancialData `bson:"financial data"`
	Earnings      requests.Earnings
}

type Locate struct {
	Address string `bson:"address,omitempty"`
	City    string `bson:"city,omitempty"`
	State   string `bson:"state,omitempty"`
	Country string `bson:"country,omitempty"`
}

type Contacts struct {
	Phone   string `bson:"phone,omitempty"`
	Website string `bson:"website,omitempty"`
}

type Staff struct {
	Employees       float64 `bson:"employees,omitempty"`
	CompanyOfficers []requests.CompanyOfficer
}

type FinancialData struct {
	TotalCash         float64 `bson:"total cash ,omitempty"`
	TotalCashPerShare float64 `bson:"total cash per share,omitempty"`
	Ebitda            float64 `bson:"ebitda,omitempty"`
	TotalDebt         float64 `bson:"total debt,omitempty"`
	QuickRatio        float64 `bson:"quick ratio,omitempty"`
	CurrentRatio      float64 `bson:"current ratio,omitempty"`
	TotalRevenue      float64 `bson:"total revenue,omitempty"`
	RevenuePerShare   float64 `bson:"revenue per share,omitempty"`
	DebtToEquity      float64 `bson:"debt to equity,omitempty"`
	ReturnOnAssets    float64 `bson:"roa,omitempty"`
	ReturnOnEquity    float64 `bson:"roe,omitempty"`
}

func main() {
	tickers, err := tickers.GetTickers()
	if err != nil {
		log.Fatal("GetTickers: ", err)
	}

	var invalidTickers sync.Map
	var errorTickers sync.Map
	var validTickers sync.Map

	wg := sync.WaitGroup{}
	wg.Add(len(tickers))
	limit := make(chan struct{}, 100)

	getData := func(ticker string) {
		defer wg.Done()
		requests.GetData(ticker, &invalidTickers, &errorTickers, &validTickers)
		<-limit
	}

	start := time.Now()

	for ticker, _ := range tickers {
		limit <- struct{}{}
		go getData(ticker)
	}

	wg.Wait()
	printRequestResults(&invalidTickers, &errorTickers, &validTickers, time.Since(start), len(tickers))

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("stock_market").Collection("stocks")

	docs := []interface{}{}

	for _, ticker := range *getKeys(&validTickers) {
		if value, ok := validTickers.Load(ticker); ok {
			data := value.(struct {
				Price         requests.Price
				AssetProfile  requests.AssetProfile
				Earnings      requests.Earnings
				FinancialData requests.FinancialData
			})

			docs = append(docs, Stock{
				Symbol:    ticker,
				ShortName: data.Price.ShortName,
				LongName:  data.Price.LongName,
				Summary:   data.AssetProfile.LongBusinessSummary,
				Industry:  data.AssetProfile.Industry,
				Sector:    data.AssetProfile.Sector,
				Staff: Staff{
					Employees:       data.AssetProfile.FullTimeEmployees,
					CompanyOfficers: data.AssetProfile.CompanyOfficers,
				},
				Locate: Locate{
					Address: data.AssetProfile.Address1,
					City:    data.AssetProfile.City,
					State:   data.AssetProfile.State,
					Country: data.AssetProfile.Country,
				},
				Contacts: Contacts{
					Phone:   data.AssetProfile.Phone,
					Website: data.AssetProfile.Website,
				},
				FinancialData: FinancialData{
					TotalCash:         float64(data.FinancialData.TotalCash),
					TotalCashPerShare: float64(data.FinancialData.TotalCashPerShare),
					Ebitda:            float64(data.FinancialData.TotalCash),
					TotalDebt:         float64(data.FinancialData.TotalDebt),
					QuickRatio:        float64(data.FinancialData.QuickRatio),
					CurrentRatio:      float64(data.FinancialData.CurrentRatio),
					TotalRevenue:      float64(data.FinancialData.TotalRevenue),
					RevenuePerShare:   float64(data.FinancialData.RevenuePerShare),
					DebtToEquity:      float64(data.FinancialData.DebtToEquity),
					ReturnOnAssets:    float64(data.FinancialData.ReturnOnAssets),
					ReturnOnEquity:    float64(data.FinancialData.ReturnOnEquity),
				},
				Earnings: data.Earnings,
			})

		}
	}

	_, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	/*ticker := "AAPL"

	requests.GetData(ticker, &invalidTickers, &errorTickers, &validTickers)

	if v, ok := validTickers.Load(ticker); ok {
		fmt.Println(v.(struct {
			Price requests.Price
			AssetProfile  requests.AssetProfile
			Earnings      requests.Earnings
			FinancialData requests.FinancialData
		}))
	}*/
}

func printRequestResults(invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map, duration time.Duration, countTickers int) {
	printSyncMapLength(invalidTickers, "Invalid Tickers count", countTickers)
	printSyncMapLength(errorTickers, "Error Tickers count", countTickers)
	printSyncMapLength(validTickers, "Valid Tickers count", countTickers)
	printScratchDuration(duration)
}

func printSyncMapLength(m *sync.Map, msg string, countTickers int) {
	count := syncMapLen(m)
	fmt.Printf("%s %d/%d\n", msg, count, countTickers)
}

func printScratchDuration(duration time.Duration) {
	duration = duration.Round(time.Second)
	fmt.Printf("Duration: %d min %.f sec\n", int(duration.Minutes()), math.Mod(duration.Seconds(), 60))
}

func syncMapLen(m *sync.Map) int {
	length := 0
	m.Range(func(key interface{}, value interface{}) bool {
		length++
		return true
	})
	return length
}

func getKeys(m *sync.Map) *[]string {
	keys := make([]string, 0)
	m.Range(func(key interface{}, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	sort.Strings(keys)
	return &keys
}
