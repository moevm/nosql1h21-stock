<template>
  <Spinner v-if="!allSectors" />
  <template v-else>
    <label><input type="checkbox" v-model="selectSector" />Select sector</label>
    <select :disabled="!selectSector" v-model="sector">
        <option v-for="sector in allSectors" :value="sector" :key="sector">{{ sector }}</option>
    </select>
    <label><input type="checkbox" :disabled="!selectSector" v-model="selectIndustry" />Select industry</label>
    <Spinner v-if="selectIndustry && !industries" />
    <select v-else :disabled="!selectSector || !selectIndustry" v-model="industry">
        <option v-for="industry in industries" :value="industry" :key="industry">{{ industry }}</option>
    </select>
  </template>
</template>

<script>
import Spinner from "@/components/Spinner";

export default {
  components: {Spinner},
  emits: ["update:sector", "update:industry"],
  data() {
    return {
      allSectors: null,
      sector: "",
      selectSector: false,
      industries: null,
      industry: "",
      selectIndustry: false,
    }
  },
  methods: {
    updateSector() {
      this.$emit("update:sector", this.selectSector ? this.sector : "")
    },
    updateIndustry() {
      this.$emit("update:industry", this.selectIndustry ? this.industry : "")
    },
  },
  watch: {
    sector() {
      this.updateSector()
      this.industries = null,
      this.updateIndustry()
      fetch('http://127.0.0.1:3000/industries-in-sector/' + this.sector)
      .then(response => response.json())
      .then(industries => {
        this.industries = industries
      })
    },
    selectSector() {
      this.updateSector()
    },
    industry() {
      this.updateIndustry()
    },
    selectIndustry() {
      this.updateIndustry()
    },
  },
  created() {
    fetch('http://127.0.0.1:3000/sectors')
    .then(response => response.json())
    .then(sectors => {
      this.allSectors = sectors
    })
  }
}
</script>

<style scoped>
  label {
    display: block;
    margin: 0 auto;
  }
</style>