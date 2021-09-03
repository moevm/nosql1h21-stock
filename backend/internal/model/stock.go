package model

type StockOverview struct {
	Symbol    string
	ShortName string `bson:"short name,omitempty"`
}

type Stock struct {
	Symbol        string `bson:"symbol"`
	ShortName     string `bson:"short name,omitempty"`
	LongName      string `bson:"long name,omitempty"`
	Summary       string `bson:"summary,omitempty"`
	Industry      string `bson:"industry,omitempty"`
	Sector        string `bson:"sector,omitempty"`
	Staff         Staff
	Locate        Locate
	Contacts      Contacts
	FinancialData FinancialData `bson:"financial data"`
	Earnings      Earnings
}

type Staff struct {
	Employees       int `bson:"employees,omitempty"`
	CompanyOfficers []CompanyOfficer
}

type CompanyOfficer struct {
	Name     string
	Age      int    `bson:"age,omitempty"`
	Title    string `bson:"title,omitempty"`
	YearBorn int    `bson:"year born,omitempty"`
	TotalPay int    `bson:"total pay,omitempty"`
}

type Locate struct {
	Address string `bson:"address,omitempty"`
	City    string `bson:"city,omitempty"`
	State   string `bson:"state,omitempty"`
	Country string `bson:"country,omitempty"`
}

type Contacts struct {
	Phone   string `bson:"phone,omitempty"`
	Website string `bson:"website,omitempty"`
}

type FinancialData struct {
	TotalCash         float64 `bson:"total cash,omitempty"`
	TotalCashPerShare float64 `bson:"total cash per share,omitempty"`
	Ebitda            float64 `bson:"ebitda,omitempty"`
	TotalDebt         float64 `bson:"total debt,omitempty"`
	QuickRatio        float64 `bson:"quick ratio,omitempty"`
	CurrentRatio      float64 `bson:"current ratio,omitempty"`
	TotalRevenue      float64 `bson:"total revenue,omitempty"`
	RevenuePerShare   float64 `bson:"revenue per share,omitempty"`
	DebtToEquity      float64 `bson:"debt to equity,omitempty"`
	ReturnOnAssets    float64 `bson:"roa,omitempty"`
	ReturnOnEquity    float64 `bson:"roe,omitempty"`
}

type Earnings struct {
	Yearly            []Year
	Quarterly         []Quarter
	FinancialCurrency string `bson:"financial currency,omitempty"`
}

type Year struct {
	Date     int
	Revenue  int
	Earnings int
}

type Quarter struct {
	Date     string
	Revenue  int
	Earnings int
}
