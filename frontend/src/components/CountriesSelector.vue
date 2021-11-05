<template>
  <Spinner v-if="!allCountries" />
  <template v-else>
    <label><input type="checkbox" v-model="selectCountries" />Select countries</label>
    <select :disabled="!selectCountries" v-model="countries" multiple>
        <option v-for="country in allCountries" :value="country" :key="country">{{ country }}</option>
    </select>
  </template>
</template>

<script>
import Spinner from "@/components/Spinner";

export default {
  components: {Spinner},
  emits: ["update:countries"],
  data() {
    return {
      allCountries: null,
      countries: [],
      selectCountries: false,
    }
  },
  methods: {
    update() {
      this.$emit("update:countries", this.selectCountries ? this.countries : [])
    },
  },
  watch: {
    countries() {
      this.update()
    },
    selectCountries() {
      this.update()
    }
  },
  created() {
    fetch('http://127.0.0.1:3000/countries')
    .then(response => response.json())
    .then(countries => {
      this.allCountries = countries
    })
  }
}
</script>

<style scoped>
  label {
    display: block;
    margin: 0 auto;
  }
  select {
    height: 100px;
  }
</style>
