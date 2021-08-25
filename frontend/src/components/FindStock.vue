<template>
  <form @submit.prevent="onSubmit">
    <input type="text">
    <button type="submit">Create</button>
  </form>
</template>

<script>
export default {

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