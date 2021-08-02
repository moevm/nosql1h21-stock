package processing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const directory = "screener_files/"

func ProccessFiles(tickers *map[string]struct{}, filenames ...string) {
	for _, filename := range filenames {
		proccessFile(filename, tickers)
	}
}

func proccessFile(filename string, tickers *map[string]struct{}) {

	file, err := os.Open(directory + filename)
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
		if _, ok := (*tickers)[ticker]; !ok {
			(*tickers)[ticker] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
