package model

type Error struct {
	Error string
}

type Stock struct {
	//ID            primitive.ObjectID `bson:"_id,omitempty"`
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

type Staff struct {
	Employees       float64 `bson:"employees,omitempty"`
	CompanyOfficers []CompanyOfficer
}

type FinancialData struct {
	TotalCash         float64 `bson:"total cash ,omitempty"`
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

type Content float64

type ValidTicker struct {
	Symbol    string `bson:"symbol"`
	ShortName string `bson:"short name"`
}

type SortRequest struct {
	Countries []string
	Industry  string
	Sector    string
}

type Sector struct {
	Sector     string
	Industries []string
}

type ValidData struct {
	Sectors []Sector
	Tickers []ValidTicker
}
