package requests

import (
	"encoding/json"
)

type Data struct {
	QuoteSummary struct {
		Result []struct {
			Price         Price
			AssetProfile  AssetProfile
			Earnings      Earnings
			FinancialData FinancialData
		}
		Error interface{}
	}
}

type Price struct {
	ShortName string
	LongName  string
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
	CompanyOfficers     []CompanyOfficer
}

type CompanyOfficer struct {
	Name     string
	Age      int     `bson:"age,omitempty"`
	Title    string  `bson:"title,omitempty"`
	YearBorn int     `bson:"year born,omitempty"`
	TotalPay Content `bson:"total pay,omitempty"`
}

type Earnings struct {
	FinancialsChart struct {
		Yearly    []Year
		Quarterly []Quarter
	}
	FinancialCurrency string
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

type contentAlias = struct {
	Raw float64
}

type Content float64

func (c *Content) UnmarshalJSON(b []byte) error {
	var alias contentAlias
	err := json.Unmarshal(b, &alias)
	if err != nil {
		err = json.Unmarshal(b, (*float64)(c))
	} else {
		*c = Content(alias.Raw)
	}
	return err
}
