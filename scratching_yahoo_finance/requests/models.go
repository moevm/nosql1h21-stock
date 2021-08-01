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
	Date    int
	Revenue struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	Earnings struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
}

type Quarter struct {
	Date    string
	Revenue struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	Earnings struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
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
	TotalCash struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	TotalCashPerShare struct {
		Raw float64
		Fmt string
	}
	Ebitda struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	TotalDebt struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	QuickRatio struct {
		Raw float64
		Fmt string
	}
	CurrentRatio struct {
		Raw float64
		Fmt string
	}
	TotalRevenue struct {
		Raw     float64
		Fmt     string
		LongFmt string
	}
	RevenuePerShare struct {
		Raw float64
		Fmt string
	}
	DebtToEquity struct {
		Raw float64
		Fmt string
	}
	ReturnOnAssets struct {
		Raw float64
		Fmt string
	}
	ReturnOnEquity struct {
		Raw float64
		Fmt string
	}
}
