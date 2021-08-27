<template>
  <p>Name: {{ stock.ShortName }}</p>
  <p>Ticker: {{ stock.Symbol }}</p>
  <p>Sector: {{ stock.Sector }}</p>
  <p>Industry: {{ stock.Industry }}</p>
  <p>Summary: {{ stock.Summary }}</p>
</template>

<script>
export default {
  name: "CompanyData",
  props: {
    stock: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>

</style>