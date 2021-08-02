package requests

import (
	"bytes"
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

func jsonToString(j []byte) string {
	buf := bytes.Buffer{}
	json.Indent(&buf, j, "", "  ")
	return buf.String()
}

func GetData(ticker string, invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map) {
	body := getBody(&ticker, setModules(profileModule, earningsModule, financialDataModule), errorTickers)

	if body == nil {
		return
	}

	data := Data{}
	err := json.Unmarshal(*body, &data)

	if err != nil {
		errorTickers.Store(ticker, struct{}{})
		log.Println("Json unmarshal error:", err, "ticker", ticker, jsonToString(*body))
		return
	}

	if data.QuoteSummary.Error != nil {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s was not found", ticker)
		return
	}

	earnings := data.QuoteSummary.Result[0].Earnings
	if earnings.FinancialCurrency == "" || len(earnings.FinancialsChart.Quarterly) == 0 ||
		len(earnings.FinancialsChart.Yearly) == 0 {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s hasn't earnings data", ticker)
		return
	}

	finData := data.QuoteSummary.Result[0].FinancialData
	if finData.TotalCash.Raw == 0 {
		invalidTickers.Store(ticker, struct{}{})
		log.Printf("Ticker %5s hasn't financial data", ticker)
		return
	}

	validTickers.Store(ticker, data.QuoteSummary.Result[0])
}

func setModules(v ...string) string {
	modules := v[0]
	for i := 1; i < len(v); i++ {
		modules = modules + "," + v[i]
	}
	return modules
}
