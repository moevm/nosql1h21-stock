<template>
  <button v-on:click="isShowFilters = !isShowFilters">{{btnFilterText()}}</button>
  <div v-show="isShowFilters">
    <div>
      <CountriesSelector v-if="disable!='country'" @update:countries="countries = $event"/>
      <SectorIndustrySelector v-if="disable!='sector' && disable!='industry'" @update:sector="sector = $event" @update:industry="industry = $event"/>
      <br>
    </div>

    <table class="content-table">
      <tbody>
      <tr>
        <td><input type="checkbox" id="Employees" value="Employees" v-model="checkedNames"></td>
        <td><label for="Employees">Employees</label></td>
        <td>
          <select v-model="employees_selected">
            <option v-for="option in options" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
        </td>
        <td><input type="number" v-model="Employees" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Total Cash" value="Total Cash" v-model="checkedNames"></td>
        <td><label for="Total Cash">Total Cash</label></td>
        <td><select v-model="total_cash_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.TotalCash" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Total Cash per share" value="Total Cash per share" v-model="checkedNames"></td>
        <td><label for="Total Cash per share">Total Cash per share</label></td>
        <td><select v-model="total_cash_per_share_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.TotalCashPerShare" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Ebitda" value="Ebitda" v-model="checkedNames"></td>
        <td><label for="Ebitda">Ebitda</label></td>
        <td><select v-model="ebitda_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.Ebitda" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Total Debt" value="Total Debt" v-model="checkedNames"></td>
        <td><label for="Total Debt">Total Debt</label></td>
        <td><select v-model="total_debt_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.TotalDebt" placeholder=""></td>
      </tr>


      <tr>
        <td><input type="checkbox" id="Quick ratio" value="Quick ratio" v-model="checkedNames"></td>
        <td><label for="Quick ratio">Quick ratio</label></td>
        <td><select v-model="quick_ratio_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.QuickRatio" placeholder=""></td>
      </tr>


      <tr>
        <td><input type="checkbox" id="Current ratio" value="Current ratio" v-model="checkedNames"></td>
        <td><label for="Current ratio">Current ratio</label></td>
        <td><select v-model="current_ratio_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.CurrentRatio" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Total Revenue" value="Total Revenue" v-model="checkedNames">
        </td>
        <td><label for="Total Revenue">Total Revenue</label></td>
        <td><select v-model="total_revenue_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.TotalRevenue" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Revenue per share" value="Revenue per share" v-model="checkedNames"></td>
        <td><label for="Revenue per share">Revenue per share</label></td>
        <td><select v-model="revenue_per_share_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.RevenuePerShare" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Debt to equity" value="Debt to equity" v-model="checkedNames"></td>
        <td><label for="Debt to equity">Debt to equity</label></td>
        <td><select v-model="debt_to_equity_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.DebtToEquity" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Roa" value="Roa" v-model="checkedNames"></td>
        <td><label for="Roa">Roa</label></td>
        <td><select v-model="roa_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.Roa" placeholder=""></td>
      </tr>

      <tr>
        <td><input type="checkbox" id="Roe" value="Roe" v-model="checkedNames"></td>
        <td><label for="Roe">Roe</label></td>
        <td><select v-model="roe_selected">
          <option v-for="option in options" v-bind:value="option.value">
            {{ option.text }}
          </option>
        </select></td>
        <td><input type="number" v-model="this.Roe" placeholder=""></td>
      </tr>

      </tbody>
    </table>
    <br>
  </div>
</template>

<script>
import CountriesSelector from "@/components/CountriesSelector";
import SectorIndustrySelector from "@/components/SectorIndustrySelector";

export default {
  components: {CountriesSelector, SectorIndustrySelector},
  props: ["disable"],
  emits: ["update:filter"],
  data() {
    return {
      isShowFilters: false,
      countries: [],
      sector: "",
      industry: "",
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
  methods: {
    update() {
      let params = []

      if (this.sector !== "") params.push("sector=" + encodeURIComponent(this.sector))
      if (this.industry !== "") params.push("industry=" + encodeURIComponent(this.industry))
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
      if (this.countries.length > 0) params.push("countries=" + this.countries.join())

      this.$emit("update:filter", params.join("&"))
    },
    getFilter(filter) {
      if (filter === ">=") {
        return ">"
      } else if (filter === "=<") {
        return "<"
      } else if (filter === "=") {
        return ""
      }
    },
    btnFilterText() {
      if (this.isShowFilters){
        return "Hide Filter"
      } else {
        return "Show Filter"
      }
    },
  },
  watch: {
    checkedNames(){
      this.update()
      this.page = 1
    },
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
  padding: 12px 25px;
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
button {
  margin-bottom: 10px;
}
</style>
