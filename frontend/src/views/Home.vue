<template>
  <h2>Search page</h2>
  <FindStock/>
  <StockList
      v-bind:tickers="tickers"
  />
</template>

<script>
import FindStock from "@/components/FindStock";
import StockList from "@/components/StockList";

export default {
  components: {StockList, FindStock},
  data() {
    return {
      countries: [],
      sectors: [],
      tickers: [],
      loading: true,
    }
  },
  mounted() {
    fetch('http://127.0.0.1:3000/validData')
        .then(response => response.json())
        .then(json => {
          setTimeout(() => {
            this.countries = json.Countries
            this.sectors = json.Sectors
            this.tickers = json.Tickers
            this.loading = false
          }, 1000)
        })
  }
}
</script>