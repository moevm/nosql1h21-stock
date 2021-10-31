<template>
  <div>
    <!--    <input type="checkbox" id="jack" value="Джек" v-model="checkedNames">-->
    <!--    <label for="jack">Джек</label>-->
    <!--    <input type="checkbox" id="jack2" value="Джек" v-model="checkedNames">-->
    <!--    <label for="jack2">Джек2</label>-->
    <!--    <input type="checkbox" id="jack3" value="Джек" v-model="checkedNames">-->
    <!--    <label for="jack3">Джек3</label>-->
    <!--    <input type="checkbox" id="jack4" value="Джек" v-model="checkedNames">-->
    <!--    <label for="jack4">Джек4</label>-->

    <input type="checkbox" id="Employees" value="Employees" v-model="checkedNames">
    <label for="Employees">Employees</label>
    <select v-model="employees_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="Employees" placeholder="">
    <br>

    <input type="checkbox" id="Total Cash" value="Total Cash" v-model="checkedNames">
    <label for="Total Cash">Total Cash</label>
    <select v-model="total_cash_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.TotalCash" placeholder="">
    <br>

    <input type="checkbox" id="Total Cash per share" value="Total Cash per share" v-model="checkedNames">
    <label for="Total Cash per share">Total Cash per share</label>
    <select v-model="total_cash_per_share_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.TotalCashPerShare" placeholder="">
    <br>

    <input type="checkbox" id="Ebitda" value="Ebitda" v-model="checkedNames">
    <label for="Ebitda">Ebitda</label>
    <select v-model="ebitda_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.Ebitda" placeholder="">
    <br>

    <input type="checkbox" id="Total Debt" value="Total Debt" v-model="checkedNames">
    <label for="Total Debt">Total Debt</label>
    <select v-model="total_debt_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.TotalDebt" placeholder="">
    <br>

    <input type="checkbox" id="Quick ratio" value="Quick ratio" v-model="checkedNames">
    <label for="Quick ratio">Quick ratio</label>
    <select v-model="quick_ratio_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.QuickRatio" placeholder="">
    <br>

    <input type="checkbox" id="Current ratio" value="Current ratio" v-model="checkedNames">
    <label for="Current ratio">Current ratio</label>
    <select v-model="current_ratio_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.CurrentRatio" placeholder="">
    <br>

    <input type="checkbox" id="Total Revenue" value="Total Revenue" v-model="checkedNames">
    <label for="Total Revenue">Total Revenue</label>
    <select v-model="total_revenue_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.TotalRevenue" placeholder="">
    <br>

    <input type="checkbox" id="Revenue per share" value="Revenue per share" v-model="checkedNames">
    <label for="Revenue per share">Revenue per share</label>
    <select v-model="revenue_per_share_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.RevenuePerShare" placeholder="">
    <br>

    <input type="checkbox" id="Debt to equity" value="Debt to equity" v-model="checkedNames">
    <label for="Debt to equity">Debt to equity</label>
    <select v-model="debt_to_equity_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.DebtToEquity" placeholder="">
    <br>

    <input type="checkbox" id="Roa" value="Roa" v-model="checkedNames">
    <label for="Roa">Roa</label>
    <select v-model="roa_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.Roa" placeholder="">
    <br>

    <input type="checkbox" id="Roe" value="Roe" v-model="checkedNames">
    <label for="Roe">Roe</label>
    <select v-model="roe_selected">
      <option v-for="option in options" v-bind:value="option.value">
        {{ option.text }}
      </option>
    </select>
    <input type="number" v-model="this.Roe" placeholder="">
    <br>

    <span>Отмеченные имена: {{ checkedNames }}</span>

    <!--    <br>
        <select v-model="employees_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select>
        <span>Выбрано: {{ employees_selected }}</span>

        <br>
        <select v-model="total_cash_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select>
        <span>Выбрано: {{ total_cash_selected }}</span>-->

  </div>


  <h2>Table</h2>
  <table class="content-table">


    <thead>
    <tr>
      <th>Ticker</th>
      <th>Name</th>
      <th>Sector</th>
      <th>Industry</th>
      <th>Country</th>
      <th v-if="this.checkedNames.indexOf(this.EmployeesString) >= 0">Employees</th>
      <th v-if="this.checkedNames.indexOf(this.TotalCashString) >= 0">Total Cash</th>
      <th v-if="this.checkedNames.indexOf(this.TotalCashPerShareString) >= 0">Total Cash per share</th>
      <th v-if="this.checkedNames.indexOf(this.EbitdaString) >= 0">Ebitda</th>
      <th v-if="this.checkedNames.indexOf(this.TotalDebtString) >= 0">Total Debt</th>
      <th v-if="this.checkedNames.indexOf(this.QuickRatioString) >= 0">Quick ratio</th>
      <th v-if="this.checkedNames.indexOf(this.CurrentRatioString) >= 0">Current ratio</th>
      <th v-if="this.checkedNames.indexOf(this.TotalRevenueString) >= 0">Total Revenue</th>
      <th v-if="this.checkedNames.indexOf(this.RevenuePerShareString) >= 0">Revenue per share</th>
      <th v-if="this.checkedNames.indexOf(this.DebtToEquityString) >= 0">Debt to equity</th>
      <th v-if="this.checkedNames.indexOf(this.RoaString) >= 0">Roa</th>
      <th v-if="this.checkedNames.indexOf(this.RoeString) >= 0">Roe</th>
    </tr>
    </thead>
    <tbody>
    <TableRow
        v-for="stock of stocks" :key="stock.Symbol"
        v-bind:stock="stock"
        v-bind:checkedNames="checkedNames"
    />
    <!--    <tr>
          <td>CLDX</td>
          <td>Celldex Therapeutics, Inc.</td>
          <td>Healthcare</td>
          <td>Biotechnology</td>
          <td>United States</td>
        </tr>
        <tr class="active-row">
          <td>2</td>
          <td>Sally</td>
          <td>72,400</td>
          <td>Students</td>
          <td>dcodeewdwefw</td>
        </tr>
        <tr>
          <td>3</td>
          <td>Nick</td>
          <td>52,300</td>
          <td>dcode</td>
          <td>dcodeewdwefw</td>
        </tr>-->
    </tbody>
  </table>

</template>

<script>

import TableRow from "@/components/TableRow";

export default {
  components: {TableRow},
  data() {
    return {
      countries: [],
      sector: "",
      industry: "",
      stocks: [],
      // employees: "",
      checkedNames: [],

      employees_selected: '>=',
      total_cash_selected: '>=',
      total_cash_per_share_selected: '>=',
      ebitda_selected: '>=',
      total_debt_selected: '>=',
      quick_ratio_selected: '>=',
      current_ratio_selected: '>=',
      total_revenue_selected: '>=',
      revenue_per_share_selected: '>=',
      debt_to_equity_selected: '>=',
      roa_selected: '>=',
      roe_selected: '>=',

      options: [
        {text: '>=', value: '>=', filter: '>'},
        {text: '=<', value: '=<', filter: '<'},
        {text: '=', value: '=', filter: ''},
      ],

      Employees: 0,
      TotalCash: 0,
      TotalCashPerShare: 0,
      Ebitda: 0,
      TotalDebt: 0,
      QuickRatio: 0,
      CurrentRatio: 0,
      TotalRevenue: 0,
      RevenuePerShare: 0,
      DebtToEquity: 0,
      Roa: 0,
      Roe: 0,
    }
  },
  created() {
    this.EmployeesString = "Employees"
    this.TotalCashString = "Total Cash"
    this.TotalCashPerShareString = "Total Cash per share"
    this.EbitdaString = "Ebitda"
    this.TotalDebtString = "Total Debt"
    this.QuickRatioString = "Quick ratio"
    this.CurrentRatioString = "Current ratio"
    this.TotalRevenueString = "Total Revenue"
    this.RevenuePerShareString = "Revenue per share"
    this.DebtToEquityString = "Debt to equity"
    this.RoaString = "Roa"
    this.RoeString = "Roe"
    this.Employees = 0
    this.TotalCash = 0
    this.TotalCashPerShare = 0
    this.Ebitda = 0
    this.TotalDebt = 0
    this.QuickRatio = 0
    this.CurrentRatio = 0
    this.TotalRevenue = 0
    this.RevenuePerShare = 0
    this.DebtToEquity = 0
    this.Roa = 0
    this.Roe = 0
    this.update()
  },
  methods: {
    update() {

      let params = []

      if (this.checkedNames.indexOf(this.EmployeesString) >= 0) params.push("employees=" + encodeURIComponent(this.getFilter(this.employees_selected) + this.Employees))
      if (this.checkedNames.indexOf(this.TotalCashString) >= 0) params.push("total cash=" + encodeURIComponent(this.getFilter(this.total_cash_selected) + this.TotalCash))
      if (this.checkedNames.indexOf(this.TotalCashPerShareString) >= 0) params.push("total cash per share=" + encodeURIComponent(this.getFilter(this.total_cash_per_share_selected) + this.TotalCashPerShare))
      if (this.checkedNames.indexOf(this.EbitdaString) >= 0) params.push("ebitda=" + encodeURIComponent(this.getFilter(this.ebitda_selected) + this.Ebitda))
      if (this.checkedNames.indexOf(this.TotalDebtString) >= 0) params.push("total debt=" + encodeURIComponent(this.getFilter(this.total_debt_selected) + this.TotalDebt))
      if (this.checkedNames.indexOf(this.QuickRatioString) >= 0) params.push("quick ratio=" + encodeURIComponent(this.getFilter(this.quick_ratio_selected) + this.QuickRatio))
      if (this.checkedNames.indexOf(this.CurrentRatioString) >= 0) params.push("current ratio=" + encodeURIComponent(this.getFilter(this.current_ratio_selected) + this.CurrentRatio))
      if (this.checkedNames.indexOf(this.TotalRevenueString) >= 0) params.push("total revenue=" + encodeURIComponent(this.getFilter(this.total_revenue_selected) + this.TotalRevenue))
      if (this.checkedNames.indexOf(this.RevenuePerShareString) >= 0) params.push("revenue per share=" + encodeURIComponent(this.getFilter(this.revenue_per_share_selected) + this.RevenuePerShare))
      if (this.checkedNames.indexOf(this.DebtToEquityString) >= 0) params.push("debt to equity=" + encodeURIComponent(this.getFilter(this.debt_to_equity_selected) + this.DebtToEquity))
      if (this.checkedNames.indexOf(this.RoaString) >= 0) params.push("roa=" + encodeURIComponent(this.getFilter(this.roa_selected) + this.Roa))
      if (this.checkedNames.indexOf(this.RoeString) >= 0) params.push("roe=" + encodeURIComponent(this.getFilter(this.roe_selected) + this.Roe))

      /*let url = "http://127.0.0.1:3000/table?" + [
        "sector=" + encodeURIComponent(this.sector),
        "industry=" + encodeURIComponent(this.industry),
        "employees=" + encodeURIComponent(this.getFilter(this.employees_selected) + this.Employees),
        "total cash=" + encodeURIComponent(this.getFilter(this.total_cash_selected) + this.TotalCash),
        "total cash per share=" + encodeURIComponent(this.getFilter(this.total_cash_per_share_selected) + this.TotalCashPerShare),
        "ebitda" + encodeURIComponent(this.getFilter(this.ebitda_selected) + this.Ebitda),
        "total debt" + encodeURIComponent(this.getFilter(this.total_debt_selected) + this.TotalDebt),
        "quick ratio" + encodeURIComponent(this.getFilter(this.quick_ratio_selected) + this.QuickRatio),
        "current ratio" + encodeURIComponent(this.getFilter(this.current_ratio_selected) + this.CurrentRatio),
        "total revenue" + encodeURIComponent(this.getFilter(this.total_revenue_selected) + this.TotalRevenue),
        "revenue per share" + encodeURIComponent(this.getFilter(this.revenue_per_share_selected) + this.RevenuePerShare),
        "debt to equity" + encodeURIComponent(this.getFilter(this.debt_to_equity_selected) + this.DebtToEquity),
        "roa" + encodeURIComponent(this.getFilter(this.roa_selected) + this.Roa),
        "roe" + encodeURIComponent(this.getFilter(this.roe_selected) + this.Roe),
        "countries=" + this.countries.join(),
      ].join("&")*/
      let url = "http://127.0.0.1:3000/table?" + params.join("&")
      fetch(url)
          .then(response => response.json())
          .then(stocks => {
            this.stocks = stocks
            this.stocks.sort(function (a, b) {
              if (a.Symbol > b.Symbol) {
                return 1;
              }
              if (a.Symbol < b.Symbol) {
                return -1;
              }
              // a должно быть равным b
              return 0;
            })
          })
      console.log(url)
      // console.log(this.getFilter(this.employees_selected))
    },
    isIn(string2) {
      return string2 in this.checkedNames
    },
    getFilter(filter) {
      if (filter === ">=") {
        return ">"
      } else if (filter === "=<") {
        return "<"
      } else if (filter === "=") {
        return ""
      }
    }
  },
  watch: {
    countries() {
      this.update()
    },
    sector() {
      this.update()
    },
    industry() {
      this.update()
    },
    employees_selected() {
      this.update()
    },
    total_cash_selected() {
      this.update()
    },
    total_cash_per_share_selected() {
      this.update()
    },
    ebitda_selected() {
      this.update()
    },
    total_debt_selected() {
      this.update()
    },
    quick_ratio_selected() {
      this.update()
    },
    current_ratio_selected() {
      this.update()
    },
    total_revenue_selected() {
      this.update()
    },
    revenue_per_share_selected() {
      this.update()
    },
    debt_to_equity_selected() {
      this.update()
    },
    roa_selected() {
      this.update()
    },
    roe_selected() {
      this.update()
    },
    Employees() {
      this.update()
    },
    TotalCash() {
      this.update()
    },
    TotalCashPerShare() {
      this.update()
    },
    Ebitda() {
      this.update()
    },
    TotalDebt() {
      this.update()
    },
    QuickRatio() {
      this.update()
    },
    CurrentRatio() {
      this.update()
    },
    TotalRevenue() {
      this.update()
    },
    RevenuePerShare() {
      this.update()
    },
    DebtToEquity() {
      this.update()
    },
    Roa() {
      this.update()
    },
    Roe() {
      this.update()
    },
  }
}
</script>

<style scoped>

.content-table {
  margin: 0 auto;
  border-collapse: collapse;
  /*margin: 25px 0;*/
  font-size: 0.9em;
  min-width: 400px;
  border-radius: 5px 5px 0 0;
  overflow: hidden;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
}

.content-table thead tr {
  background-color: #009879;
  color: #ffffff;
  text-align: left;
  font-weight: bold;
}

.content-table th,
.content-table td {
  padding: 12px 15px;
}

.content-table tbody tr {
  border-bottom: 1px solid #dddddd;
}

.content-table tbody tr:nth-of-type(even) {
  background-color: #f3f3f3;
}

.content-table tbody tr:last-of-type {
  border-bottom: 2px solid #009879;
}

.content-table tbody tr.active-row {
  font-weight: bold;
  color: #009879;
}

</style>