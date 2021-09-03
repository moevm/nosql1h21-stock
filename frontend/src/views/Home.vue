<template>
  <h2>Search page</h2>
  <input v-model.trim="searchText">
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
      searchText: "",
      pending: 0,
      stocks: [],
    }
  },
  watch: {
    searchText(searchText) {
      if (searchText === "") {
        this.stocks = []
        return
      }
      searchText = searchText.toUpperCase()
      this.pending++
      fetch('http://127.0.0.1:3000/search-by-ticker/' + searchText)
      .then(response => response.json())
      .then(stocks => {
        this.stocks = stocks
      })
      .finally(() => this.pending--)
    }
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