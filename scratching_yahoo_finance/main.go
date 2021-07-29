package main

import (
	"bufio"
	"fmt"
	"log"
	"nosql1h21-stock/scratching_yahoo_finance/requests"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	tickers := make(map[string]struct{})

	file, err := os.Open("scratching_yahoo_finance/screener_file/nasdaq_screener_1626544763604.csv")
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
	var validTickers sync.Map

	wg := sync.WaitGroup{}
	wg.Add(len(tickers))

	getInfo := func(ticker string) {
		defer wg.Done()
		requests.GetProfile(ticker, &invalidTickers, &errorTickers, &validTickers)
	}

	start := time.Now()

	count := 0
	for ticker, _ := range tickers {
		if count%100 == 0 {
			time.Sleep(4 * time.Second)
		}
		go getInfo(ticker)
		count++
	}

	wg.Wait()
	duration := time.Since(start)

	printSyncMapLength(&invalidTickers, "Invalid Tickers count", len(tickers))
	printSyncMapLength(&errorTickers, "Error Tickers count", len(tickers))
	printSyncMapLength(&validTickers, "Valid Tickers count", len(tickers))
	printScratchDuration(duration)

	/*requests.GetEarnings("AAPL", &invalidTickers, &errorTickers, &validTickers)

	if v , ok := validTickers.Load("AAPL"); ok {
		fmt.Println(v)
	}*/
}

func printSyncMapLength(m *sync.Map, msg string, countTickers int) {
	count := 0
	m.Range(func(key interface{}, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("%s %d/%d\n", msg, count, countTickers)
}

func printScratchDuration(duration time.Duration) {
	duration = duration.Round(time.Second)
	seconds := int(duration.Seconds()) % 60
	fmt.Printf("Duration %.f min %d sec \n", duration.Minutes(), seconds)
}
