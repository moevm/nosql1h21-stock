package tickers

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

const directory = "internal/scratcher/screener_files/"

func GetTickers() (map[string]struct{}, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	tickers := map[string]struct{}{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		err = getTickers(directory+entry.Name(), tickers)
		if err != nil {
			return nil, err
		}
	}
	return tickers, nil
}

func getTickers(filepath string, tickers map[string]struct{}) error {
	if !strings.HasSuffix(filepath, ".csv") {
		return nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // Skipping the header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) > 0 {
			ticker := record[0]
			ticker = strings.Split(ticker, "^")[0]
			ticker = strings.Split(ticker, "/")[0]
			ticker = strings.TrimSpace(ticker)
			tickers[ticker] = struct{}{}
		}
	}

	return nil
}
