<script setup>

import ParticleEffect from "../components/ParticleEffect.vue";
</script>

<template>
  <ParticleEffect />

  <div class="cont d-flex flex-column justify-content-center align-content-center align-items-center noselect">
    <h1>Simple<span class="badge text-wrap">GoLang</span>chat</h1>
    <div class="my-2">
      <router-link to="/register" class="btn btn-outline-light mx-2">Register</router-link>
      <router-link to="/sign" class="btn btn-outline-danger">Log in</router-link>
    </div>
  </div>
</template>

<style scoped>
.cont {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  z-index: 3;
}

.noselect {
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>