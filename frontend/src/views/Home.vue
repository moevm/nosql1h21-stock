<template>
  <h2>Search page</h2>
  <input v-model.trim="fragment">
  <label><input type="checkbox" v-model="selectCountries" />Select countries</label>
  <select :disabled="!selectCountries" v-model="countries" multiple>
    <option v-for="country in allCountries" :value="country" :key="country">{{ country }}</option>
  </select>
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

export default {
  components: {Spinner },
  data() {
    return {
      fragment: "",
      allCountries: [],
      countries: [],
      selectCountries: false,
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
        "countries=" + (this.selectCountries ? this.countries.join() : "")
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
    }
  },
  created() {
    this.pending++
    fetch('http://127.0.0.1:3000/countries')
    .then(response => response.json())
    .then(countries => {
      this.allCountries = countries
    })
    .finally(() => this.pending--)
  }
}
</script>

<style scoped>
  label {
    display: block;
    margin: 0 auto;
  }
  select {
    height: 300px;
  }
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