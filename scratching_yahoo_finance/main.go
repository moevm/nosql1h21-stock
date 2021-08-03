package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"nosql1h21-stock/scratching_yahoo_finance/processing"
	"nosql1h21-stock/scratching_yahoo_finance/requests"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Stock struct {
	Symbol        string  `bson:"symbol"`
	ShortName     string  `bson:"short name,omitempty"`
	LongName      string  `bson:"long name,omitempty"`
	Summary       string  `bson:"summary,omitempty"`
	Industry      string  `bson:"industry,omitempty"`
	Sector        string  `bson:"sector,omitempty"`
	Employees     float64 `bson:"employees,omitempty"`
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

	tickers := make(map[string]struct{})

	processing.ProccessFiles(&tickers, "nasdaq_screener_1626544763604.csv", "constituents_csv.csv")

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

	validTickers.Range(func(key interface{}, value interface{}) bool {

		data := value.(struct {
			Price         requests.Price
			AssetProfile  requests.AssetProfile
			Earnings      requests.Earnings
			FinancialData requests.FinancialData
		})

		docs = append(docs, Stock{
			Symbol:    key.(string),
			ShortName: data.Price.ShortName,
			LongName:  data.Price.LongName,
			Summary:   data.AssetProfile.LongBusinessSummary,
			Employees: data.AssetProfile.FullTimeEmployees,
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
				TotalCash:         data.FinancialData.TotalCash.Raw,
				TotalCashPerShare: data.FinancialData.TotalCashPerShare.Raw,
				Ebitda:            data.FinancialData.TotalCash.Raw,
				TotalDebt:         data.FinancialData.TotalDebt.Raw,
				QuickRatio:        data.FinancialData.QuickRatio.Raw,
				CurrentRatio:      data.FinancialData.CurrentRatio.Raw,
				TotalRevenue:      data.FinancialData.TotalRevenue.Raw,
				RevenuePerShare:   data.FinancialData.RevenuePerShare.Raw,
				DebtToEquity:      data.FinancialData.DebtToEquity.Raw,
				ReturnOnAssets:    data.FinancialData.ReturnOnAssets.Raw,
				ReturnOnEquity:    data.FinancialData.ReturnOnEquity.Raw,
			},
			Earnings: data.Earnings,
		})
		return true
	})

	_, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	/*err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}*/
	/*databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)*/

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
