<template>
  <h2>Search page</h2>
  <FindStock
      v-on:search-text="updateSearchText"
  />
  <StockList
      v-bind:tickers="filteredTickers"
  />
</template>

<script>
import FindStock from "@/components/FindStock";
import StockList from "@/components/StockList";

export default {
  components: {StockList, FindStock},
  data() {
    return {
      searchText: "",
      countries: [],
      sectors: [],
      tickers: [],
      loading: true,
    }
  },
  methods: {
    updateSearchText(searchText) {
      this.searchText = searchText
    },
  },
  computed: {
    filteredTickers() {
      let filterTickers = []

      if (this.searchText.trim() !== "") {
        filterTickers = this.tickers.filter(t => t.Symbol === this.searchText.toUpperCase())

        if (filterTickers.length === 0) {
          filterTickers = this.tickers.filter(t => t.ShortName.toLowerCase().indexOf(this.searchText.toLowerCase()) !== -1)
        }
      }

      return filterTickers
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