<template>
  <h2>Search page</h2>
  <input v-model.trim="fragment">
  <CountriesSelector @update:countries="countries = $event"/>
  <SectorIndustrySelector @update:sector="sector = $event" @update:industry="industry = $event" />
  <div class="results">Results: {{stocks.length}}</div>
  <Spinner v-if="pending"/>
  <div v-else class="stock" v-for="stock of stocks" :key="stock.Symbol" @click="this.$router.push('/stock/' + stock.Symbol)">
    <div>{{stock.Symbol}}</div>
    <div class="pusher" />
    <div>{{stock.ShortName}}</div>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";
import CountriesSelector from "@/components/CountriesSelector";
import SectorIndustrySelector from "@/components/SectorIndustrySelector";

export default {
  components: {Spinner, CountriesSelector, SectorIndustrySelector},
  data() {
    return {
      fragment: "",
      countries: [],
      sector: "",
      industry: "",
      pending: 0,
      stocks: [],
    }
  },
  methods: {
    update() {
      this.fragment = this.fragment.toLowerCase()
      this.pending++
      let url = "http://127.0.0.1:3000/search?" + [
        "fragment=" + this.fragment,
        "countries=" + this.countries.join(),
        "sector=" + encodeURIComponent(this.sector),
        "industry=" + encodeURIComponent(this.industry),
      ].join("&")
      fetch(url)
      .then(response => response.json())
      .then(stocks => {
        this.stocks = stocks
      })
      .finally(() => this.pending--)
    },
  },
  watch: {
    fragment() {
      this.update()
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
  }
}
</script>

<style scoped>
  .stock {
    width: 600px;
    margin: 0 auto;
    padding: 4px 0;
    border-top: 1px solid #ccc;
    display: flex;
  }
  .stock:hover {
    cursor: pointer;
    background-color: #f5f5f5;
  }
  .results {
    padding: 5px 0;
  }
  .pusher {
    flex-grow: 1;
  }
</style>