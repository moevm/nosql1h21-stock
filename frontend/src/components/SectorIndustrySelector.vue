<template>
  <Spinner v-if="!allSectors" />
  <template v-else>
    <label><input type="checkbox" v-model="selectSector" />Select sector</label>
    <select :disabled="!selectSector" v-model="sector">
        <option v-for="sector in allSectors" :value="sector" :key="sector">{{ sector }}</option>
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
      industry: "",
      selectIndustry: false,
    }
  },
  methods: {
    updateSector() {
      this.$emit("update:sector", this.selectSector ? this.sector : "")
    },
  },
  watch: {
    sector() {
      this.updateSector()
    },
    selectSector() {
      this.updateSector()
    }
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