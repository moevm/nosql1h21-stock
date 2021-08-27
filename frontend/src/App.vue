<template>
  <div id="app">
  <h1>Stock application</h1>
  <nav>
    <router-link to="/">Search</router-link>
  </nav>
  <hr>
  <router-view/>
  </div>
</template>

<style scoped>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>