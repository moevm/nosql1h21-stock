package scratcher

import (
	"context"
	"log"
	"sync"
	"time"

	"nosql1h21-stock-backend/backend/internal/scratcher/mongodb"
	"nosql1h21-stock-backend/backend/internal/scratcher/tickers"
	"nosql1h21-stock-backend/backend/internal/scratcher/yahoo"

	"go.mongodb.org/mongo-driver/mongo"
)

func Scratch(ctx context.Context, collection *mongo.Collection) {
	tickers, err := tickers.GetTickers()
	if err != nil {
		log.Fatal("GetTickers: ", err)
	}

	mu := sync.Mutex{}
	companiesInfo := map[string]*yahoo.CompanyInfo{}

	wg := sync.WaitGroup{}
	limit := make(chan struct{}, 100)

	start := time.Now()
	log.Println("Getting companies info")

	for ticker := range tickers {
		wg.Add(1)
		limit <- struct{}{}
		ticker := ticker
		go func() {
			defer func() {
				wg.Done()
				<-limit
			}()
			companyInfo, err := yahoo.GetCompanyInfo(ticker)
			if err != nil {
				log.Printf("Getting info about %v failed: %v\n", ticker, err)
				return
			}
			incompleteInfo := companyInfo.Price.CompanyLongName == "" ||
				companyInfo.Price.CompanyShortName == "" ||
				companyInfo.AssetProfile == nil ||
				companyInfo.AssetProfile.Industry == "" ||
				companyInfo.AssetProfile.Sector == "" ||
				companyInfo.Earnings == nil ||
				companyInfo.FinancialData == nil ||
				companyInfo.FinancialData.TotalCash == 0 ||
				companyInfo.Earnings.FinancialCurrency == ""
				// companyInfo.AssetProfile.FullTimeEmployees == 0
				// companyInfo.FinancialData.QuickRatio == 0 ||
				// companyInfo.FinancialData.CurrentRatio == 0 ||
				// companyInfo.FinancialData.DebtToEquity == 0
			if incompleteInfo {
				log.Printf("Company with ticker %v excluded because of incomplete info\n", ticker)
				return
			}
			mu.Lock()
			defer mu.Unlock()
			companiesInfo[ticker] = companyInfo
		}()
	}
	wg.Wait()

	log.Println("Full info was got for", len(companiesInfo), "of", len(tickers), "tickers in", time.Since(start))

	start = time.Now()
	log.Println("Saving the info to the database")

	err = mongodb.SaveCompaniesInfo(collection, companiesInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved in", time.Since(start))
}
