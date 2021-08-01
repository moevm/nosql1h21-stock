package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const profileModule = "assetProfile"
const earningsModule = "earnings"
const financialDataModule = "financialData"

func getBody(ticker *string, module string, errorTickers *sync.Map) *[]byte {
	client := http.Client{
		Timeout: time.Duration(time.Minute),
	}

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v10/finance/quoteSummary/%s?modules=%s", *ticker, module)

	res, err := client.Get(url)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Get error:", err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Body error:", err, "ticker", ticker)
		return nil
	}
	return &body
}

func GetProfile(ticker string, invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map) {
	body := getBody(&ticker, profileModule, errorTickers)

	if body == nil {
		return
	}

	profile := Profile{}

	err := json.Unmarshal(*body, &profile)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Json unmarshal error:", err, "ticker", ticker)
		return
	}

	if profile.QuoteSummary.Error != nil {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s was not found", ticker)
		return
	}

	validTickers.Store(ticker, profile.QuoteSummary.Result[0].AssetProfile)
}

func GetEarnings(ticker string, invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map) {
	body := getBody(&ticker, earningsModule, errorTickers)

	if body == nil {
		return
	}

	earnings := Earnings{}
	err := json.Unmarshal(*body, &earnings)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Json unmarshal error:", err, "ticker", ticker)
		return
	}

	if earnings.QuoteSummary.Error != nil {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s was not found", ticker)
		return
	}

	validTickers.Store(ticker, earnings.QuoteSummary.Result[0].Earnings)
}

func GetFinancialData(ticker string, invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map) {
	body := getBody(&ticker, financialDataModule, errorTickers)

	if body == nil {
		return
	}

	finData := Financials{}
	err := json.Unmarshal(*body, &finData)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Json unmarshal error:", err, "ticker", ticker)
		return
	}

	if finData.QuoteSummary.Error != nil {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s was not found", ticker)
		return
	}

	validTickers.Store(ticker, finData.QuoteSummary.Result[0].FinancialData)
}
