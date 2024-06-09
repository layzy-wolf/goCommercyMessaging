<script setup>

import {inject, onBeforeMount, ref, watch} from "vue";
import {router} from "../router.js";
import * as jwt from "jsonwebtoken-esm";
import axios from "axios";
import {useRoute} from "vue-router";

const chats = ref()
const updateChat = inject("updateChat")
const apiConfig = inject("apiConfig")
const openModalAddContact = inject("openModalAddContact")
const openModalCreateGroup = inject("openModalCreateGroup")
const openModalEnterGroup = inject("openModalEnterGroup")
const route = useRoute()

console.log(route.matched)

onBeforeMount(async () => {
  await update()
})

let user = jwt.decode(localStorage.getItem("token"))

async function update() {
  let token = localStorage.getItem("token")

  await axios.post(apiConfig.http + "/list", {}, {
    headers: {
      Authorization: token
    }
  }).then(res => {
    if (res.data.users !== undefined) {
      let arr = {}
      for (let v of res.data.users) {
        arr[v] = {}
      }
      chats.value = arr
    }
  }).catch(err => {
    console.log(err)
  })

  let ts = localStorage.getItem("timestamp")

  if (ts === undefined || ts === null) {
    ts = "2024-01-01 10:00:00"
  }

  await axios.post(apiConfig.http + "/chats", {
      timestamp: ts
    }, {
      headers: {
        Authorization: token
      }
    })
        .then(res => {
          let today = new Date();
          let dd = String(today.getDate()).padStart(2, '0')
          let mm = String(today.getMonth() + 1).padStart(2, '0')
          let yyyy = today.getFullYear()
          today = yyyy + '-' + mm + '-' + dd + " " +
              String(today.getHours()).padStart(2, '0') + ":" +
              String(today.getMinutes()).padStart(2, '0') + ":" +
              String(today.getSeconds()).padStart(2, '0')

          localStorage.setItem("timestamp", today.toString())
          let d = chats.value

          for (let key in d) {
            d[key].new = false
          }

          if (res.data.messages !== null) {
            res.data.messages.forEach(m => {
              let login = jwt.decode(token).login
              if (jwt.decode(m.messageHead).from !== login) {
                for (let key in d) {
                  if (key === jwt.decode(m.messageHead).from) {
                    delete d[key]
                    d[key] = {message: jwt.decode(m.messageBody).message, new: true}
                    console.log(key)
                  }
                }
              }
            })
            chats.value = d
          }
        })
}

watch(updateChat, () => {
  update()
})

function openAddContact() {
  openModalAddContact.value = true
}

function openCreateGroup() {
  openModalCreateGroup.value = true
}

function enterGroup() {
  openModalEnterGroup.value = true
}

function deleteAccount() {
  axios.post(apiConfig.http + "/remove", {}, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  }).then(res => {
    if (res.data.success) {
      logout()
    }
  }).catch()
}

function logout() {
  localStorage.removeItem("token")
  router.push("/sign")
}

</script>

<template>
  <div class="cont border-start">
    <div class="d-flex justify-content-between align-items-center px-2">
      <img src="/Logo.svg" alt="Logo" height="48">
      <div class="dropdown">
        <a href="#"
           class="d-flex align-items-center justify-content-center p-3 link-body-emphasis text-decoration-none dropdown-toggle"
           data-bs-toggle="dropdown" aria-expanded="false">
          <h5>@</h5>
        </a>
        <ul class="dropdown-menu text-small shadow" style="">
          <li class="dropdown-header border-bottom">
            <h5 class="text-light">{{ user.login }}</h5>
          </li>
          <li><a class="dropdown-item" href="#" @click.prevent="openAddContact">Добавить Контакт</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="openCreateGroup">Создать Группу</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="enterGroup">Войти в группу</a></li>
          <li><hr class="dropdown-divider"></li>
          <li><a class="dropdown-item" href="#" @click.prevent="deleteAccount">Удалить Аккаунт</a></li>
          <li><hr class="dropdown-divider"></li>
          <li><a class="dropdown-item" href="#" @click.prevent="logout">Выйти</a></li>
        </ul>
      </div>
    </div>
    <div class="d-flex flex-column align-items-stretch flex-shrink-0">
      <div class="list-group list-group-flush border-bottom custom-scroll overflow-y-auto h-100">
        <router-link v-for="(chat, user) in chats" :to="'/chat/' + user" class="list-group-item list-group-item-action py-3 lh-sm"
                     aria-current="true">
          <div class="d-flex w-100 align-items-center justify-content-between h-100">
            <strong class="mb-1">{{ user }}</strong>
            <small v-if="chat.new">Новое</small>
          </div>
          <div class="col-10 mb-1 small">{{ chat.message }}</div>
        </router-link>

      </div>
    </div>
  </div>

</template>

<style scoped>
.cont {
  box-sizing: border-box;
  height: 100vh;
  width: 25vw;
  overflow: hidden;
}

.custom-scroll {
  overflow-y: scroll;
  max-height: 100vh !important;
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 1) rgba(0, 0, 0, 0);
  transition: scrollbar-color .3s ease;
  -webkit-overflow-scrolling: touch;
  pointer-events: auto;
}

.list-group-item {
  background-color: rgba(0, 0, 0, 0);
}

.list-group-item.active {
  background-color: rgba(0, 0, 0, .5);
  border-color: rgba(0, 0, 0, 0);
}
@media (width < 800px) {
  .responsive {
    width: 100%;
  }
}
</style>