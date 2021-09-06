package yahoo

import (
	"encoding/json"
)

type Number float64

func (c *Number) UnmarshalJSON(b []byte) error {
	var complex struct {
		Raw float64
	}
	err := json.Unmarshal(b, &complex)
	if err != nil {
		err = json.Unmarshal(b, (*float64)(c))
	} else {
		*c = Number(complex.Raw)
	}
	return err
}

type Response struct {
	QuoteSummary struct {
		Result []CompanyInfo
		Error  json.RawMessage
	}
}

type CompanyInfo struct {
	Price         Price
	AssetProfile  *AssetProfile
	Earnings      *Earnings
	FinancialData *FinancialData
}

type Price struct {
	// There are many fields, but we only need the company name
	CompanyShortName string `json:"shortName"`
	CompanyLongName  string `json:"longName"`
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
	FullTimeEmployees   int
	CompanyOfficers     []CompanyOfficer
}

type CompanyOfficer struct {
	Name     string
	Age      int
	Title    string
	YearBorn int
	TotalPay Number
}

type Earnings struct {
	FinancialsChart struct {
		Yearly    []Year
		Quarterly []Quarter
	}
	FinancialCurrency string
}

type Year struct {
	Date     int
	Revenue  Number
	Earnings Number
}

type Quarter struct {
	Date     string
	Revenue  Number
	Earnings Number
}

type FinancialData struct {
	TotalCash         Number
	TotalCashPerShare Number
	Ebitda            Number
	TotalDebt         Number
	QuickRatio        Number
	CurrentRatio      Number
	TotalRevenue      Number
	RevenuePerShare   Number
	DebtToEquity      Number
	ReturnOnAssets    Number
	ReturnOnEquity    Number
}
