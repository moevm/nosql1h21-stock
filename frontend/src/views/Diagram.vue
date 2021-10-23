<template>
  <h2>Diagram</h2>
  <Spinner v-if="!countItems"/>
  <div v-else class="root">
    <p>Count by
      <select v-model="countBy">
        <option value="country">Country</option>
        <option value="industry">Industry</option>
        <option value="sector">Sector</option>
      </select>
    </p>
    <div class="diagram">
      <div class="item" v-for="item in countItems" :key="item.Key">
        <div class="label">{{item.Key}}</div>
        <div class="line" :style="{width: (item.Amount / maxAmount * 300) + 'px'}" />
        <div>{{item.Amount}}</div>
      </div>
    </div>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";

export default {
  components: {Spinner},
  data() {
    return {
      countItems: null,
      countBy: "",
    }
  },
  created() {
    this.countBy = "country"
  },
  computed: {
    maxAmount() {
      let max = 0
      for (let item of this.countItems) {
        max = Math.max(max, item.Amount)
      }
      return max
    },
  },
  watch: {
    countBy() {
        this.countItems = null
        fetch('http://127.0.0.1:3000/count?by=' + this.countBy)
        .then(response => response.json())
        .then(countItems => {
            this.countItems = countItems
        })
    }
  },
}
</script>

<style scoped>
  .root {
    margin: 0 auto;
    max-width: 600px;
  }
  .diagram {
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    gap: 5px;
  }
  .item {
    display: inline-flex;
    align-items: center;
    gap: 10px;
  }
  .label {
    line-height: 1;
    width: 200px;
  }
  .line {
    display: inline-block;
    min-width: 1px;
    height: 10px;
    background-color: green;
  }
</style>
