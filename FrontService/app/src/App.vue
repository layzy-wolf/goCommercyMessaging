<script setup>

import LeftBar from "./components/LeftBar.vue";
</script>

<template>
  <div class="main d-flex bg-dark" data-bs-theme="dark">
    <router-view></router-view>
  </div>
</template>

<style scoped>
.main {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  height: 100vh;
}
</style>

<style>
  :root {
    --bs-dark-rgb: 24, 24, 24 !important;
    --bs-body-bg: #181818 !important;
    --bs-dark: #181818 !important;
    --bs-tertiary-bg-rgb: 43, 48, 53 !important;
    --bs-tertiary-bg: #181818 !important;
    --bs-list-group-bg: #181818 !important;
  }
</style>