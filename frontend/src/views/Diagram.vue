<template>
  <h2>Diagram</h2>
  <div class="root">
    <p>Mode
      <select v-model="mode">
        <option value="count">Count companies</option>
        <option value="average">Average</option>
      </select>
    </p>

    <p v-if="mode=='average'">Property
      <select v-model="property">
        <option value="employees">Employees amount</option>
        <!-- <option value="total cash">Total cash</option> -->
        <!-- <option value="total cash per share">Total cash per share</option> -->
        <!-- <option value="ebitda">Ebitda</option> -->
        <!-- <option value="total debt">Total debt</option> -->
        <option value="quick ratio">Quick ratio</option>
        <option value="current ratio">Current ratio</option>
        <!-- <option value="total revenue">Total revenue</option> -->
        <!-- <option value="revenue per share">Revenue per share</option> -->
        <option value="debt to equity">Debt to equity</option>
        <option value="roa">Roa</option>
        <option value="roe">Roe</option>
      </select>
    </p>
    <p>In
      <select v-model="countBy">
        <option value="country">Country</option>
        <option value="industry">Industry</option>
        <option value="sector">Sector</option>
      </select>
    </p>
    <Filter @update:filter="filter = $event" :disable='countBy'/>
    <Spinner v-if="!countItems"/>
    <div v-else class="diagram">
      <div class="item" v-for="item in countItems" :key="item.Key">
        <div class="label">{{item.Key}}</div>
        <div class="line" :style="{width: (item.Amount / maxAmount * 300) + 'px'}" />
        <div>{{Math.round(item.Amount * 1000) / 1000 + item.Unit}}</div>
        <!-- <div>{{item.Amount + item.Unit}}</div> -->
      </div>
    </div>
  </div>
</template>

<script>
import Spinner from "@/components/Spinner";
import Filter from "@/components/Filter";

export default {
  components: {Spinner, Filter},
  data() {
    return {
      mode: "count",
      countItems: null,
      countBy: "country",
      filter: "",
      property: "employees",
    }
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
  created() {
    this.update()
  },
  watch: {
    mode() {
      this.update()
    },
    countBy() {
      this.update()
    },
    filter() {
      this.update()
    },
    property() {
      this.update()
    },
  },
  methods: {
    update() {
        this.countItems = null
        let property = this.mode === "average" ? "&property=" + this.property : ""
        let filter = this.filter ? '&' + this.filter : ""
        fetch('http://127.0.0.1:3000/aggregate?mode=' + this.mode + property + '&in=' + this.countBy + filter)
        .then(response => response.json())
        .then(countItems => {
            this.countItems = countItems
        })
    },
  }
}
</script>

<style scoped>
  .root {
    margin: 0 auto;
    max-width: 700px;
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
