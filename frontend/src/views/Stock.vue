<template>
  <h2>Stock</h2>
  <Spinner v-if="!stock"/>
  <div v-else class="stock">
    <p>Name: {{ stock.ShortName }}</p>
    <p>Ticker: {{ stock.Symbol }}</p>
    <p>Sector: {{ stock.Sector }}</p>
    <p>Industry: {{ stock.Industry }}</p>
    <p>Summary: {{ stock.Summary }}</p>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";

export default {
  components: {Spinner},
  data() {
    return {
      stock: null,
    }
  },
  created() {
    let ticker = this.$route.params.ticker
    fetch('http://127.0.0.1:3000/stock/' + ticker)
    .then(response => response.json())
    .then(stock => {
        this.stock = stock
    })
  }
}
</script>