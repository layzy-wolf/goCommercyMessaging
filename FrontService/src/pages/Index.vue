<script setup>

import ParticleEffect from "../components/ParticleEffect.vue";
import axios from "axios";
import {inject, onBeforeMount} from "vue";
import {router} from "../router.js";

const apiConfig = inject(["apiConfig"])

onBeforeMount(() => {
  let token = localStorage.getItem("token")

  if (token !== null) {
    axios.post(apiConfig.http + "/verify", {}, {
      headers: {
        Authorization: token
      }
    })
        .then(res => {
          if (res.data.valid) {
            router.push("/chat")
          } else {
            localStorage.removeItem("token")
          }
        })
        .catch(err => {
          localStorage.removeItem("token")
    })
  }
})

</script>

<template>
  <ParticleEffect />

  <div class="cont d-flex flex-column justify-content-center align-content-center align-items-center noselect">
    <h1>Простой<span class="badge text-wrap">GoLang</span>Чат</h1>
    <div class="my-2">
      <router-link to="/register" class="btn btn-outline-light mx-2">Регистрация</router-link>
      <router-link to="/sign" class="btn btn-outline-danger">Войти</router-link>
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