<template>
  <h2>Comparator</h2>
  <div v-if="toCompare().length < 2">
    Add at least 2 stocks to compare.
  </div>
  <div v-else class="stock">
    <table>
      <tr v-for="parameter in parameters" :key="parameter.label">
        <td>{{parameter.label}}</td>
        <td :class="[compare(parameter.field, parameter.better, 0)]">{{parameter.field(toCompare()[0])}}</td>
        <td :class="[compare(parameter.field, parameter.better, 1)]">{{parameter.field(toCompare()[1])}}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import {toCompare} from "../to_compare";
import {currency} from "../currency";

export default {
  components: {},
  data() {
    let cur = (stock, amount) => {
      return currency(amount, stock.FinancialData.FinancialCurrency)
    }
    return {
      parameters: [
        {label: "Ticker", field: stock => stock.Symbol, better: 0},
        {label: "Name", field: stock => stock.ShortName, better: 0},
        {label: "Sector", field: stock => stock.Sector, better: 0},
        {label: "Industry", field: stock => stock.Industry, better: 0},
        {label: "Country", field: stock => stock.Locate.Country, better: 0},
        {label: "Total cash", field: stock => cur(stock, stock.FinancialData.TotalCash), better: 0},
        {label: "Total cash per share", field: stock => cur(stock,  stock.FinancialData.TotalCashPerShare), better: 0},
        {label: "Ebitda", field: stock => cur(stock, stock.FinancialData.Ebitda), better: 0},
        {label: "Total debt", field: stock => cur(stock, stock.FinancialData.TotalDebt), better: 0},
        {label: "Quick ratio", field: stock => stock.FinancialData.QuickRatio, better: 1},
        {label: "Current ratio", field: stock => stock.FinancialData.CurrentRatio, better: 1},
        {label: "Total revenue", field: stock => cur(stock, stock.FinancialData.TotalRevenue), better: 0},
        {label: "Revenue per share", field: stock => cur(stock, stock.FinancialData.RevenuePerShare), better: 0},
        {label: "Debt to equity", field: stock => stock.FinancialData.DebtToEquity, better: -1},
        {label: "Return on assets", field: stock => stock.FinancialData.ReturnOnAssets, better: 1},
        {label: "Return on equity", field: stock => stock.FinancialData.ReturnOnEquity, better: 1},
      ],
    }
  },
  methods: {
    toCompare() {
      return toCompare
    },
    compare(field, better, iAm) {
      if (better === 0)
        return null
      if (iAm === 1)
        better = -better
      return [
        better === 1 ? "red" : "green",
        better === 1 ? "green" : "red",
        "",
      ][field(toCompare[0]) < field(toCompare[1]) ? 0 :
        field(toCompare[0]) > field(toCompare[1]) ? 1 :
        2]
    }
  }
}
</script>

<style scoped>
  .stock {
    margin: 0 auto;
    max-width: 600px;
  }
  .red {
    background-color: #f005;
  }
  .green {
    background-color: #0f05;
  }

</style>