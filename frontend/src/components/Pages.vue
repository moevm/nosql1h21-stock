<template>
  <button v-if="page>1" v-on:click="$emit('prevClick')">Previous</button>
  <button v-if="page < totalPage" v-on:click="$emit('nextClick')">Next</button>
  <label>Page = {{ page }}; TotalPages = {{ totalPage }}</label>
</template>

<script>
export default {
  name: "Pages",
  emits: ['prevClick', 'nextClick'],
  props: {
    page: Number,
    totalPage: Number,
  },
}
</script>

<style scoped>

label {
  margin-left: 15px;
  margin-right: 15px;
}

button {
  margin-left: 15px;
  margin-right: 15px;
}

</style>