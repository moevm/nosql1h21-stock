package requests

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
	TotalCashPerShare ReducedContent
	Ebitda            Content
	TotalDebt         Content
	QuickRatio        ReducedContent
	CurrentRatio      ReducedContent
	TotalRevenue      Content
	RevenuePerShare   ReducedContent
	DebtToEquity      ReducedContent
	ReturnOnAssets    ReducedContent
	ReturnOnEquity    ReducedContent
}

type Content struct {
	Raw     float64
	Fmt     string
	LongFmt string
}

type ReducedContent struct {
	Raw float64
	Fmt string
}
