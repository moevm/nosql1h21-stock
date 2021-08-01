package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"nosql1h21-stock/scratching_yahoo_finance/requests"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	tickers := make(map[string]struct{})

	file, err := os.Open("screener_file/nasdaq_screener_1626544763604.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		ticker := strings.Split(scanner.Text(), ",")[0]
		ticker = strings.Split(ticker, "^")[0]
		ticker = strings.Split(ticker, "/")[0]
		ticker = strings.Trim(ticker, " \n")
		if _, ok := tickers[ticker]; !ok {
			tickers[ticker] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var invalidTickers sync.Map
	var errorTickers sync.Map
	var tickerProfiles sync.Map
	var tickerEarnings sync.Map
	var tickerFinancialData sync.Map

	wg := sync.WaitGroup{}
	wg.Add(len(tickers))
	limit := make(chan struct{}, 100)

	getInfo := func(ticker string) {
		defer wg.Done()
		requests.GetProfile(ticker, &invalidTickers, &errorTickers, &tickerProfiles)
		<-limit
	}

	getEarnings := func(ticker string) {
		defer wg.Done()
		requests.GetEarnings(ticker, &invalidTickers, &errorTickers, &tickerEarnings)
		<-limit
	}

	getFinancialData := func(ticker string) {
		defer wg.Done()
		requests.GetFinancialData(ticker, &invalidTickers, &errorTickers, &tickerFinancialData)
		<-limit
	}

	getResponseForMap(&invalidTickers, &errorTickers, &tickerProfiles, &tickers, &wg, getInfo, len(tickers), limit)
	wg = sync.WaitGroup{}
	getResponse(&invalidTickers, &errorTickers, &tickerEarnings, &tickerProfiles, &wg, getEarnings, len(tickers), limit)
	wg = sync.WaitGroup{}
	getResponse(&invalidTickers, &errorTickers, &tickerFinancialData, &tickerEarnings, &wg, getFinancialData, len(tickers), limit)

	/*ticker:= "AAPL"

	requests.GetFinancialData(ticker, &invalidTickers, &errorTickers, &tickerFinancialData)

	if v , ok := tickerFinancialData.Load(ticker); ok {
		fmt.Println(v.(requests.FinancialData).TotalCash)
	}*/
}

func getResponse(invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map, iterateMap *sync.Map, wg *sync.WaitGroup, doRequests func(ticket string), tickerCount int, limit chan struct{}) {
	start := time.Now()
	count := syncMapLen(iterateMap)

	wg.Add(count)

	iterateMap.Range(func(key interface{}, value interface{}) bool {
		limit <- struct{}{}
		go doRequests(key.(string))
		return true
	})

	wg.Wait()
	printRequestResults(invalidTickers, errorTickers, validTickers, time.Since(start), tickerCount)
}

func getResponseForMap(invalidTickers *sync.Map, errorTickers *sync.Map, validTickers *sync.Map, tickers *map[string]struct{}, wg *sync.WaitGroup, doRequests func(ticket string), tickerCount int, limit chan struct{}) {
	start := time.Now()

	for ticker, _ := range *tickers {
		limit <- struct{}{}
		go doRequests(ticker)
	}

	wg.Wait()
	printRequestResults(invalidTickers, errorTickers, validTickers, time.Since(start), tickerCount)
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
