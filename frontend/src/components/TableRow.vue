<template>
  <tr>
    <td>{{ stock.Symbol }}</td>
    <td>{{ stock.Name }}</td>
    <td>{{ stock.Sector }}</td>
    <td>{{ stock.Industry }}</td>
    <td>{{ stock.Locate.Country }}</td>
    <th v-if="this.checkedNames.indexOf(this.Employees) >= 0">{{ stock.Staff.Employees }}</th>
    <th v-if="this.checkedNames.indexOf(this.TotalCash) >= 0">{{ this.currency(stock.FinancialData.TotalCash,stock.FinancialData.FinancialCurrency) }}</th>
    <th v-if="this.checkedNames.indexOf(this.TotalCashPerShare) >= 0">{{ stock.FinancialData.TotalCashPerShare }}</th>
    <th v-if="this.checkedNames.indexOf(this.Ebitda) >= 0">{{ stock.FinancialData.Ebitda }}</th>
    <th v-if="this.checkedNames.indexOf(this.TotalDebt) >= 0">{{ stock.FinancialData.TotalDebt }}</th>
    <th v-if="this.checkedNames.indexOf(this.QuickRatio) >= 0">{{ stock.FinancialData.QuickRatio }}</th>
    <th v-if="this.checkedNames.indexOf(this.CurrentRatio) >= 0">{{ stock.FinancialData.CurrentRatio }}</th>
    <th v-if="this.checkedNames.indexOf(this.TotalRevenue) >= 0">{{ stock.FinancialData.TotalRevenue }}</th>
    <th v-if="this.checkedNames.indexOf(this.RevenuePerShare) >= 0">{{ stock.FinancialData.RevenuePerShare }}</th>
    <th v-if="this.checkedNames.indexOf(this.DebtToEquity) >= 0">{{ stock.FinancialData.DebtToEquity }}</th>
    <th v-if="this.checkedNames.indexOf(this.Roa) >= 0">{{ stock.FinancialData.ReturnOnAssets }}</th>
    <th v-if="this.checkedNames.indexOf(this.Roe) >= 0">{{ stock.FinancialData.ReturnOnEquity }}</th>
  </tr>
</template>

<script>

export default {
  props: {
    stock: {
      type: Object,
      required: true
    },
    checkedNames: {
      type: Object,
      required: true
    }
  },
  created() {
    this.Employees = "Employees"
    this.TotalCash = "Total Cash"
    this.TotalCash = "Total Cash"
    this.TotalCashPerShare = "Total Cash per share"
    this.Ebitda = "Ebitda"
    this.TotalDebt = "Total Debt"
    this.QuickRatio = "Quick ratio"
    this.CurrentRatio = "Current ratio"
    this.TotalRevenue = "Total Revenue"
    this.RevenuePerShare = "Revenue per share"
    this.DebtToEquity = "Debt to equity"
    this.Roa = "Roa"
    this.Roe = "Roe"
  },
  methods: {

    currency(amount, name) {
      let factors = ["K", "M", "B", "T"]
      let factor = ""
      for (let i = 0; i < factors.length; i++) {
        if (Math.abs(amount) < 10000)
          break
        amount = Math.round(amount)
        amount /= 1000
        factor = factors[i]
      }
      return amount + factor + " " + name
    }
  }
}

</script>
