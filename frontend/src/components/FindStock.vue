<template>
  <form @submit.prevent="onSubmit">
    <input v-model="text">
    <button @click="clickSearchBtn()">Edit</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      text: '',
      fieldDisable: false
    }
  },
  methods: {
    clickSearchBtn() {
      this.$emit('search-text', this.text)
    },
  },
  watch: {
    text: function (value) {
     // console.log('Text edited: ', value);
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