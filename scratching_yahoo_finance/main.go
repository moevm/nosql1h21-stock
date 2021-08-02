package main

import (
	"fmt"
	"math"
	"nosql1h21-stock/scratching_yahoo_finance/processing"
	"nosql1h21-stock/scratching_yahoo_finance/requests"
	"sync"
	"time"
)

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

	/*ticker := "AAPL"

	requests.GetData(ticker, &invalidTickers, &errorTickers, &validTickers)

	if v, ok := validTickers.Load(ticker); ok {
		fmt.Println(v.(struct {
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
