<template>
  <h3>{{ticker.ShortName}}</h3>
  <h4>{{ticker.Symbol}}</h4>
</template>

<script>
export default {
  name: "StockItem",
  props: {
    ticker: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>

</style>