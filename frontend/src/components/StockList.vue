<template>
  <ul>
    <StockItem
      v-for="ticker of tickers" :key="ticker.Symbol"
      v-bind:ticker="ticker"
    />
  </ul>
</template>

<script>
import StockItem from "@/components/StockItem";

export default {
  name: "StockList",
  components: {StockItem},
  props: ['tickers'],
}
</script>

<style scoped>

</style>