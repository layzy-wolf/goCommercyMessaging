<script setup>

import * as jwt from "jsonwebtoken-esm"
import axios from "axios";
import {router} from "../router.js";
import {inject, onBeforeMount} from "vue";

let login, password

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
        .catch(() => {
          localStorage.removeItem("token")
        })
  }
})


function submit(login, password) {
  if (login === undefined || password === undefined) {
    alert("Please fill all fields")
    return
  }

  let token = jwt.sign({login: "@" + login, password: password}, "my-Secret", {algorithm: "HS256"})

  axios.post(apiConfig.http + "/login", {
    token
  })
      .then(res => {
        if (res.data.token !== "") {
          localStorage.setItem("token", res.data.token)
          router.push("/chat")
        } else {
          alert(res.data.error)
        }
      })
      .catch(err => alert(err.response.data))
}

</script>

<template>

  <div class="cont d-flex flex-column justify-content-center align-content-center align-items-center">
    <router-link to="/register" class="btn btn-outline-danger my-2 ms-25">Регистрация?</router-link>
    <form action="#" @submit.prevent="submit(login, password)" class="responsive">
      <div class="form-floating mb-3 text-light">
        <input type="text" class="form-control bg-dark bg-opacity-10" id="name" placeholder="Имя пользователя" v-model="login">
        <label for="name">Имя пользователя</label>
      </div>
      <div class="form-floating text-light bg-dark bg-opacity-10">
        <input type="password" class="form-control bg-dark bg-opacity-10" id="password" placeholder="Пароль" v-model="password">
        <label for="password">Пароль</label>
      </div>
      <button type="submit" class="btn btn-outline-light form-control my-5">Войти</button>
    </form>

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

.ms-25 {
  margin-left: 20%;
}

.responsive {
  width: 25%;
}

@media (width < 800px) {
  .responsive {
    width: 80vw;
  }
}
</style>