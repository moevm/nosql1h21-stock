package requests

import "encoding/json"

type Profile struct {
	QuoteSummary struct {
		Result []struct {
			AssetProfile AssetProfile
		}
		Error interface{}
	}
}

type AssetProfile struct {
	Address1            string
	City                string
	State               string
	Country             string
	Phone               string
	Website             string
	Industry            string
	Sector              string
	LongBusinessSummary string
	FullTimeEmployees   float64
}

type Year struct {
	Date     int
	Revenue  Content
	Earnings Content
}

type Quarter struct {
	Date     string
	Revenue  Content
	Earnings Content
}

type Earnings struct {
	QuoteSummary struct {
		Result []struct {
			Earnings struct {
				FinancialsChart struct {
					Yearly    []Year
					Quarterly []Quarter
				}
				FinancialCurrency string
			}
		}
		Error interface{}
	}
}

type Financials struct {
	QuoteSummary struct {
		Result []struct {
			FinancialData FinancialData
		}
		Error interface{}
	}
}

type FinancialData struct {
	TotalCash         Content
	TotalCashPerShare Content
	Ebitda            Content
	TotalDebt         Content
	QuickRatio        Content
	CurrentRatio      Content
	TotalRevenue      Content
	RevenuePerShare   Content
	DebtToEquity      Content
	ReturnOnAssets    Content
	ReturnOnEquity    Content
}

type contentAlias = struct {
	Raw float64
}
type Content contentAlias

func (c *Content) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, (*contentAlias)(c))
	if err != nil {
		err = json.Unmarshal(b, &c.Raw)
	}
	return err
}
