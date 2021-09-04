<template>
  <h2>Stock</h2>
  <Spinner v-if="!stock"/>
  <div v-else class="stock">
    <p>Ticker: {{ stock.Symbol }}</p>
    <p>Name: {{ stock.ShortName }}</p>
    <p>Sector: {{ stock.Sector }}</p>
    <p>Industry: {{ stock.Industry }}</p>
    <p>Summary: {{ stock.Summary }}</p>
    <p>Country: {{ stock.Locate.Country }}</p>
    <p>Total cash: {{stock.FinancialData.TotalCash}}</p>
    <p>Total cash per share: {{stock.FinancialData.TotalCashPerShare}}</p>
    <p>Ebitda: {{stock.FinancialData.Ebitda}}</p>
    <p>Total debt: {{stock.FinancialData.TotalDebt}}</p>
    <p>Quick ratio: {{stock.FinancialData.QuickRatio}}</p>
    <p>Current ratio: {{stock.FinancialData.CurrentRatio}}</p>
    <p>Total revenue: {{stock.FinancialData.TotalRevenue}}</p>
    <p>Revenue per share: {{stock.FinancialData.RevenuePerShare}}</p>
    <p>Debt to equity: {{stock.FinancialData.DebtToEquity}}</p>
    <p>Return on assets: {{stock.FinancialData.ReturnOnAssets}}</p>
    <p>Return on equity: {{stock.FinancialData.ReturnOnEquity}}</p>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";

export default {
  components: {Spinner},
  data() {
    return {
      stock: null,
    }
  },
  created() {
    let ticker = this.$route.params.ticker
    fetch('http://127.0.0.1:3000/stock/' + ticker)
    .then(response => response.json())
    .then(stock => {
        this.stock = stock
    })
  }
}
</script>