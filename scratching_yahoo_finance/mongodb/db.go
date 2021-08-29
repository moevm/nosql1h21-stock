package mongodb

import (
	"context"
	"nosql1h21-stock/scratching_yahoo_finance/yahoo"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (_ *mongo.Client, disconnect func(), _ error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	return client, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}, nil
}

func SaveCompaniesInfo(client *mongo.Client, companiesInfo map[string]*yahoo.CompanyInfo) error {
	collection := client.Database("stock_market").Collection("stocks")

	documents := []interface{}{}

	for ticker, info := range companiesInfo {
		document := Stock{
			Symbol:    ticker,
			ShortName: info.Price.CompanyShortName,
			LongName:  info.Price.CompanyLongName,
			Summary:   info.AssetProfile.LongBusinessSummary,
			Industry:  info.AssetProfile.Industry,
			Sector:    info.AssetProfile.Sector,
			Staff: Staff{
				Employees:       info.AssetProfile.FullTimeEmployees,
				CompanyOfficers: nil,
			},
			Locate: Locate{
				Address: info.AssetProfile.Address1,
				City:    info.AssetProfile.City,
				State:   info.AssetProfile.State,
				Country: info.AssetProfile.Country,
			},
			Contacts: Contacts{
				Phone:   info.AssetProfile.Phone,
				Website: info.AssetProfile.Website,
			},
			FinancialData: FinancialData{
				TotalCash:         float64(info.FinancialData.TotalCash),
				TotalCashPerShare: float64(info.FinancialData.TotalCashPerShare),
				Ebitda:            float64(info.FinancialData.TotalCash),
				TotalDebt:         float64(info.FinancialData.TotalDebt),
				QuickRatio:        float64(info.FinancialData.QuickRatio),
				CurrentRatio:      float64(info.FinancialData.CurrentRatio),
				TotalRevenue:      float64(info.FinancialData.TotalRevenue),
				RevenuePerShare:   float64(info.FinancialData.RevenuePerShare),
				DebtToEquity:      float64(info.FinancialData.DebtToEquity),
				ReturnOnAssets:    float64(info.FinancialData.ReturnOnAssets),
				ReturnOnEquity:    float64(info.FinancialData.ReturnOnEquity),
			},
			Earnings: Earnings{
				Yearly:            nil,
				Quarterly:         nil,
				FinancialCurrency: info.Earnings.FinancialCurrency,
			},
		}
		for _, officer := range info.AssetProfile.CompanyOfficers {
			document.Staff.CompanyOfficers = append(document.Staff.CompanyOfficers, CompanyOfficer{
				Name:     officer.Name,
				Age:      officer.Age,
				Title:    officer.Title,
				YearBorn: officer.YearBorn,
				TotalPay: int(officer.TotalPay),
			})
		}
		for _, year := range info.Earnings.FinancialsChart.Yearly {
			document.Earnings.Yearly = append(document.Earnings.Yearly, Year{
				Date:     year.Date,
				Revenue:  int(year.Revenue),
				Earnings: int(year.Earnings),
			})
		}
		for _, quarter := range info.Earnings.FinancialsChart.Quarterly {
			document.Earnings.Quarterly = append(document.Earnings.Quarterly, Quarter{
				Date:     quarter.Date,
				Revenue:  int(quarter.Revenue),
				Earnings: int(quarter.Earnings),
			})
		}
		documents = append(documents, document)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := collection.InsertMany(ctx, documents)
	return err
}
