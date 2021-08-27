<template>
  <h2>Stock</h2>
  <CompanyData v-bind:stock="stock" v-if="!loading"/>
  <Loader v-if="loading"/>
</template>

<script>
import Loader from "@/components/Loader";
import CompanyData from "@/components/CompanyData";

export default {
  components: {CompanyData, Loader},
  data() {
    return {
      stock: Object,
      loading: true,
    }
  },
  mounted() {
    let ticker = this.$route.params.ticker
    fetch('http://127.0.0.1:3000/stock/' + ticker)
        .then(response => response.json())
        .then(json => {
          setTimeout(() => {
            this.stock = json;
            this.loading = false
          }, 1000)
        })
  }
}
</script>