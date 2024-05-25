<script setup>

import * as jwt from "jsonwebtoken-esm"
import axios from "axios";
import {router} from "../router.js";

let login, password

function submit(login, password) {
  if (login === undefined || password === undefined) {
    alert("Please fill all fields")
    return
  }

  let token = jwt.sign({login: login, password: password}, "my-Secret", {algorithm: "HS256"})

  axios.post("http://localhost:8080/login", {
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
    <router-link to="/register" class="btn btn-outline-danger my-2 ms-25">Register?</router-link>
    <form action="#" @submit.prevent="submit(login, password)" class="w-25">
      <div class="form-floating mb-3 text-light">
        <input type="text" class="form-control bg-dark bg-opacity-10" id="name" placeholder="UserName" v-model="login">
        <label for="name">UserName</label>
      </div>
      <div class="form-floating text-light bg-dark bg-opacity-10">
        <input type="password" class="form-control bg-dark bg-opacity-10" id="password" placeholder="Password" v-model="password">
        <label for="password">Password</label>
      </div>
      <button type="submit" class="btn btn-outline-light form-control my-5">Send</button>
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
</style>