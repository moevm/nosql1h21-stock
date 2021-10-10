<template>
  <div id="app">
  <h1>Stock application</h1>
  <nav>
    <div><router-link to="/">Search</router-link></div>
    <div><router-link to="/comparator">Compare</router-link></div>
    <div><router-link to="/diagram">Diagram</router-link></div>
    <div> <a href="#">Table</a></div>
    <div><a href="http://127.0.0.1:3000/export">Export data</a></div>
    <div><a href="#" @click="Import">Import data</a></div>
    <input type='file' accept='.json' id="fileInput">
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

<script>
export default {
  methods: {
    Import() {
      let fileInput = document.getElementById("fileInput")
      fileInput.onchange = () => {
        let file = fileInput.files ? fileInput.files[0] : null
        if (!file)
          return
        let reader = new FileReader()
        reader.onload = () => {
          let content = reader.result
          fetch("http://127.0.0.1:3000/import", {
            method: "POST",
            body: content
          })
          .then(() => location.reload())
        }
        reader.readAsText(file)
      }
      fileInput.click()
    }
  }
}
</script>

<style scoped>
#app {
  font-family: sans-serif;
  text-align: center;
  --color: #222;
  color: var(--color);
  margin-top: 60px;
}
input[type="file"] {
  display: none;
}
</style>
