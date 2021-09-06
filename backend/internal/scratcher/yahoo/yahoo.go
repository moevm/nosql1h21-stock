package yahoo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func doRequest(ticker string) (body []byte, err error) {
	client := http.Client{
		Timeout: time.Duration(time.Minute),
	}

	modules := "price,assetProfile,earnings,financialData"
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v10/finance/quoteSummary/%s?modules=%s", ticker, modules)

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	return
}

func jsonToString(j []byte) string {
	buf := bytes.Buffer{}
	json.Indent(&buf, j, "", "  ")
	return buf.String()
}

func GetCompanyInfo(ticker string) (*CompanyInfo, error) {
	body, err := doRequest(ticker)
	if err != nil {
		return nil, err
	}

	resp := Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("%w (ticker: %v, json: %v)", err, ticker, jsonToString(body))
	}

	if err := resp.QuoteSummary.Error; !bytes.Equal(err, []byte("null")) {
		return nil, fmt.Errorf("error in response: %v", jsonToString([]byte(resp.QuoteSummary.Error)))
	}

	return &resp.QuoteSummary.Result[0], nil
}
