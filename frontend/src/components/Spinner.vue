<template>
  <div class="container">
    <div class="spinner" />
  </div>
</template>

<style scoped>
  .container {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .spinner {
    display: inline-block;
    width: 56px;
    height: 56px;
  }
  .spinner:after {
    content: "";
    display: block;
    width: 46px;
    height: 46px;
    border-radius: 50%;
    border-width: 5px;
    border-style: solid;
    border-color: var(--color) transparent;
    animation: spinning 1.2s linear infinite;
  }
  @keyframes spinning {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>