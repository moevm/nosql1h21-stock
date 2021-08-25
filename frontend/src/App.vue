<template>
  <div id="app">
  <h1>Todo application</h1>
  <nav>
    <router-link to="/">Home</router-link>
    <router-link to="/todos">About</router-link>
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