<template>
  <div id="app">
  <h1>Stock application</h1>
  <nav>
    <router-link to="/">Search</router-link>
  </nav>
  <hr>

  <!-- Magic, https://stackoverflow.com/a/67682535/13828753 -->
  <router-view v-slot="{ Component }">
    <keep-alive include="Home">
      <component :is="Component" />
    </keep-alive>
  </router-view>
  </div>
</template>

<style scoped>
#app {
  font-family: sans-serif;
  text-align: center;
  --color: #222;
  color: var(--color);
  margin-top: 60px;
}
</style>