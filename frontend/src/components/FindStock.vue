<template>
  <div>
    <input v-model="text">
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: ''
    }
  },
  methods: {

  },
  watch: {
    text: function (value) {
      this.$emit('search-text', value)
    },
  }
}
</script>

<style scoped>
form {
  display: flex;
}

input {
  width: 400px;
}
</style>