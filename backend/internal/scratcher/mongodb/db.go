package mongodb

import (
	"context"
	"time"

	"nosql1h21-stock-backend/backend/internal/scratcher/yahoo"

	"go.mongodb.org/mongo-driver/mongo"
)

func SaveCompaniesInfo(collection *mongo.Collection, companiesInfo map[string]*yahoo.CompanyInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.Drop(ctx)

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
				FinancialCurrency: info.FinancialData.FinancialCurrency,
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

	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := collection.InsertMany(ctx, documents)
	return err
}
