<template>
  <h2>Stock</h2>
  <Spinner v-if="!stock"/>
  <div v-else class="stock">
    <button @click="addToCompare">Add to compare</button>
    <p>Ticker: {{ stock.Symbol }}</p>
    <p>Name: {{ stock.ShortName }}</p>
    <p>Sector: {{ stock.Sector }}</p>
    <p>Industry: {{ stock.Industry }}</p>
    <p>Summary: {{ stock.Summary }}</p>
    <p>Country: {{ stock.Locate.Country }}</p>
    <p>Total cash: {{cur(stock.FinancialData.TotalCash)}}</p>
    <p>Total cash per share: {{cur(stock.FinancialData.TotalCashPerShare)}}</p>
    <p>Ebitda: {{cur(stock.FinancialData.Ebitda)}}</p>
    <p>Total debt: {{cur(stock.FinancialData.TotalDebt)}}</p>
    <p>Quick ratio: {{stock.FinancialData.QuickRatio}}</p>
    <p>Current ratio: {{stock.FinancialData.CurrentRatio}}</p>
    <p>Total revenue: {{cur(stock.FinancialData.TotalRevenue)}}</p>
    <p>Revenue per share: {{cur(stock.FinancialData.RevenuePerShare)}}</p>
    <p>Debt to equity: {{stock.FinancialData.DebtToEquity}}</p>
    <p>Return on assets: {{stock.FinancialData.ReturnOnAssets}}</p>
    <p>Return on equity: {{stock.FinancialData.ReturnOnEquity}}</p>
    <p>Revenue and earnings: 
      <select v-model="earningsMode">
        <option>Yearly</option>
        <option>Quarterly</option>
      </select>
    </p>
    <div class="earningsStat">
      <template v-for="item in stock.Earnings[earningsMode]" :key="item.Date">
        <div class="date">
          <div class="date-label">{{item.Date}}</div>
          <div class="items">
            <div class="item">
              <div class="revenue" :style="{width: (item.Revenue / maxEarning * 300) + 'px'}" />
              <div>{{cur(item.Revenue)}}</div>
            </div>
            <div class="item">
              <div :class="{earnings: true, negative: item.Earnings < 0}" :style="{width: (Math.abs(item.Earnings) / maxEarning * 300) + 'px'}" />
              <div>{{cur(item.Earnings)}}</div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";
import {addToCompare} from "../to_compare";
import {currency} from "../currency";

export default {
  components: {Spinner},
  data() {
    return {
      stock: null,
      earningsMode: "Yearly",
    }
  },
  created() {
    let ticker = this.$route.params.ticker
    fetch('http://127.0.0.1:3000/stock/' + ticker)
    .then(response => response.json())
    .then(stock => {
        this.stock = stock
    })
  },
  computed: {
    maxEarning() {
      let max = 0
      for (let item of this.stock.Earnings[this.earningsMode]) {
        max = Math.max(max, item.Revenue, Math.abs(item.Earnings))
      }
      return max
    },
  },
  methods: {
    addToCompare() {
      addToCompare(this.stock)
    },
    cur(amount) {
      return currency(amount, this.stock.FinancialData.FinancialCurrency)
    }
  },
}
</script>

<style scoped>
  .stock {
    margin: 0 auto;
    max-width: 600px;
  }
  .date, .item {
    display: inline-flex;
    align-items: center;
    gap: 10px;
  }
  .earningsStat {
    display: flex;
    flex-direction: column;
    margin: 0 auto;
  }
  .revenue, .earnings {
    display: inline-block;
    height: 10px;
    background-color: green;
  }
  .items {
    display: flex;
    flex-direction: column;
  }
  .negative {
    background-color: red;
  }
</style>