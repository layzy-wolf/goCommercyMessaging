<script setup>

import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import Message from "./Message.vue";
import * as jwt from "jsonwebtoken-esm"
import {router} from "../router.js";
import {useRoute} from "vue-router";
import {inject, onBeforeMount, onMounted, ref, watch} from "vue";
import axios from "axios";

let route = useRoute()
const apiConfig = inject("apiConfig")
const updateChat = inject("updateChat")

const props = defineProps(["conn"])
const messages = ref([])
const m = inject("message")
const changed = inject("changed")
const openModalGetMembers = inject("openModalGetMembers")
const openModalGetCode = inject("openModalGetCode")
let limit = 1
let length = 0
let lastMessage
let isGroup = false

onBeforeMount(async () => {
  if (Array.from(route.params.user)[0] === "&") {
    isGroup = true
  }

  let token = localStorage.getItem("token")

  await axios.post(apiConfig.http + "/show", {
    to: route.params.user,
    limit: 20 * limit
  }, {
    headers: {
      Authorization: token
    }
  })
      .then(res => {
        if (res.data.messages !== null) {
          length = res.data.messages.length
          res.data.messages.forEach(m => {
            let type, login, mH, mB
            login = jwt.decode(token).login
            mH = jwt.decode(m.messageHead)
            mB = jwt.decode(m.messageBody)
            if (mH.from === login) {
              type = "sent"
            } else {
              type = "received"
            }
            messages.value.unshift({
              type: type,
              user: mH.from,
              content: mB.message,
              timestamp: mB.timestamp
            })
          })
        }
      })
})

let messageInp

function send(message) {
  let m = {
    head: {
      from: jwt.decode(localStorage.getItem("token")).login,
      to: route.params.user
    },
    body: {
      message: message,
      timestamp: ""
    }
  }

  messageInp.value = ""

  props.conn.send(jwt.sign(m, "my-Secret", {algorithm: "HS256"}))

  let today = new Date();
  today = today.getFullYear() + '-' + String(today.getMonth() + 1).padStart(2, '0') + '-' + String(today.getDate()).padStart(2, '0') + " " +
      String(today.getHours()).padStart(2, '0') + ":" +
      String(today.getMinutes()).padStart(2, '0') + ":" +
      String(today.getSeconds()).padStart(2, '0')

  lastMessage = {
    type: "sent",
    user: m.head.from,
    content: m.body.message,
    timestamp: today
  }

  messages.value.push(lastMessage)

  document.querySelector("#scroll").scrollTop = 9999
}

watch(changed, () => {

  let today = new Date();
  today = today.getFullYear() + '-' + String(today.getMonth() + 1).padStart(2, '0') + '-' + String(today.getDate()).padStart(2, '0') + " " +
      String(today.getHours()).padStart(2, '0') + ":" +
      String(today.getMinutes()).padStart(2, '0') + ":" +
      String(today.getSeconds()).padStart(2, '0')
  messages.value.push({
    type: "received",
    user: m.value.from,
    content: m.value.message,
    timestamp: today
  })

  setTimeout(() =>
      document.querySelector("#scroll").scrollTop = 9999, 50)
}, {deep: true})

async function scroll(event) {
  let token = localStorage.getItem("token")
  let scrollHeight = event.target.scrollHeight
  if (event.target.scrollTop < 120) {
    limit++
    await axios.post(apiConfig.http + "/show", {
      to: route.params.user,
      limit: 20 * limit
    }, {
      headers: {
        Authorization: token
      }
    })
        .then(res => {
          if (res.data.messages !== null) {
            if (length !== res.data.messages.length) {
              length = res.data.messages.length
              res.data.messages.forEach(m => {
                let type, login, mH, mB
                login = jwt.decode(token).login
                mH = jwt.decode(m.messageHead)
                mB = jwt.decode(m.messageBody)
                if (mH.from === login) {
                  type = "sent"
                } else {
                  type = "received"
                }
                messages.value.unshift({
                  type: type,
                  user: mH.from,
                  content: mB.message,
                  timestamp: mB.timestamp
                })
                event.target.scrollTop = event.target.scrollHeight - scrollHeight / 2
              })
            }
          }
        })
  }
}

function getMembers() {
  openModalGetMembers.value = true
}

function getCode() {
  openModalGetCode.value = true
}

function removeContact() {
  axios.post(apiConfig.http + "/remove/contact", {
    chat: route.params.user
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  }).then(res => {
    if (res.data.success) {
      updateChat.value = !updateChat.value
      router.push("/chat")
    }
  })
}

onMounted(() => {
  setTimeout(() => {
    document.querySelector("#scroll").scrollTop = 9999
    document.querySelector("#scroll").addEventListener("scroll", scroll)

  }, 200)
})

</script>

<template>
  <div class="cont responsive">
    <div class="cont-wrapper border-end border-bottom d-flex flex-column justify-content-end">
      <div class="h5 text-light opacity-75 py-2 border-bottom d-flex justify-content-between">
        <div>
          {{ route.params.user }}
        </div>
        <div class="dropdown z-3">
          <a href="#"
             class="d-flex align-items-center link-body-emphasis text-decoration-none px-3"
             data-bs-toggle="dropdown" aria-expanded="false">
            <h5><FontAwesomeIcon icon="fa-solid fa-ellipsis" /></h5>
          </a>
          <ul class="dropdown-menu text-small shadow" style="">
            <li><a class="dropdown-item" href="#" @click.prevent="removeContact">Удалить Контакт</a></li>
            <li v-if="isGroup"><hr class="dropdown-divider"></li>
            <li v-if="isGroup"><a class="dropdown-item" href="#" @click.prevent="getMembers">Члены Группы</a></li>
            <li v-if="isGroup"><a class="dropdown-item" href="#" @click.prevent="getCode">Получить Код</a></li>
          </ul>
        </div>
      </div>

      <div class="custom-scroll overflow-y-auto h-100" id="scroll">
        <Message v-for="message in messages" :message="message"/>
      </div>

    </div>
    <form action="#" autocomplete="off"
          class="message-zone text-muted d-flex justify-content-start align-items-center p-3 bg-dark"
          @submit.prevent="send(messageInp.value)">
      <input type="text" class="form-control form-control-lg" id="exampleFormControlInput1" placeholder="Сообщение"
             ref="messageInp">
      <button class="ms-3 btn" type="submit">
        <FontAwesomeIcon :icon="['fas', 'paper-plane']"/>
      </button>
    </form>
  </div>

</template>

<style scoped>
.cont {
  padding: 0;
  box-sizing: border-box;
  height: 100vh;
  width: 60%;
  overflow: hidden;
  z-index: 3;
  margin: 0 5% 0 5%;
}

.cont-wrapper {
  height: 90vh;
  grid-template-columns: min-content min-content;
  border-image: linear-gradient(rgba(0, 0, 0, 0), #495057) 30;
}

.message-zone {
  height: 10vh;
}

.custom-scroll {
  overflow-x: hidden;
  overflow-y: scroll;
  max-height: 100vh !important;
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 1) rgba(0, 0, 0, 0);
  transition: scrollbar-color .3s ease;
  -webkit-overflow-scrolling: touch;
  pointer-events: auto;
}

@media (width < 800px) {
  .responsive {
    width: 100%;
  }
}
</style>