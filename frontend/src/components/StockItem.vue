<template>
  <div @click="clickOnStock()">
    <h3>{{ ticker.ShortName }}</h3>
    <h4>{{ ticker.Symbol }}</h4>
  </div>
</template>

<script>
export default {
  name: "StockItem",
  props: {
    ticker: {
      type: Object,
      required: true
    }
  },
  methods: {
    clickOnStock() {
      this.$router.push({ name: 'stock', params: { ticker: this.ticker.Symbol } })
    }
  }
}
</script>

<style scoped>

</style>