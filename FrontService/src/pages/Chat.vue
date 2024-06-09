<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {computed, onBeforeUnmount, ref, watch, provide, getCurrentInstance, onBeforeMount, inject} from "vue";
import {router} from "../router.js";
import {onBeforeRouteUpdate, useRoute} from "vue-router";
import ChatBar from "../components/ChatBar.vue";
import ChatWindow from "../components/ChatWindow.vue";
import * as jwt from "jsonwebtoken-esm"
import ChatNotifictaion from "../components/ChatNotifictaion.vue";
import axios from "axios";
import AddContactModal from "../components/AddContactModal.vue";
import CreateGroupModal from "../components/CreateGroupModal.vue";
import EnterGroupModal from "../components/EnterGroupModal.vue";
import GetMembersModal from "../components/GetMembersModal.vue";
import GetCodeModal from "../components/GetCodeModal.vue";
const route = useRoute()

const apiConfig = inject("apiConfig")


let conn = new WebSocket(apiConfig.ws + "/chat?Authorization=" + localStorage.getItem("token"))
let message = ref()
let modal = ref(false)
let changed = ref(false)
let updateChat = ref(false)
let openModalAddContact = ref(false)
let openModalCreateGroup = ref(false)
let openModalEnterGroup = ref(false)
let openModalGetMembers = ref(false)
let openModalGetCode = ref(false)

provide("message", message)
provide("changed", changed)
provide("modal", modal)
provide("updateChat", updateChat)
provide("openModalAddContact", openModalAddContact)
provide("openModalCreateGroup", openModalCreateGroup)
provide("openModalEnterGroup", openModalEnterGroup)
provide("openModalGetMembers", openModalGetMembers)
provide("openModalGetCode", openModalGetCode)

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
          } else {
            localStorage.removeItem("token")
            router.push("/sign")
          }
        })
        .catch(() => {
          localStorage.removeItem("token")
          router.push("/sign")
        })
  }
})

let user = jwt.decode(localStorage.getItem("token"))

conn.onmessage = function (event) {
  let m = jwt.decode(event.data)
  let mH = jwt.decode(m.MessageHead)
  let mB = jwt.decode(m.MessageBody)

  if (mH.from !== user.login) {
    message.value = {from: mH.from, message: mB.message, timestamp: mB.timestamp}

    if (route.params.user === mH.from) {
      changed.value = !changed.value
    } else {
      modal.value = !modal.value
    }

    let c = JSON.parse(localStorage.getItem("chats"))

    c[message.value.from] = {message: message.value.message, new: true}

    localStorage.setItem("chats", JSON.stringify(c))
    updateChat.value = !updateChat.value
  }
}

onBeforeUnmount(() => {
  conn.close()
})

watch(() => route.params.user, (newU, oldU) => {
  route.params.user = newU
})
</script>

<template>
  <GetMembersModal v-show="openModalGetMembers" />
  <GetCodeModal v-show="openModalGetCode" />
  <EnterGroupModal v-show="openModalEnterGroup" />
  <CreateGroupModal v-show="openModalCreateGroup" />
  <AddContactModal v-show="openModalAddContact" />
  <ChatBar />
  <ChatWindow :conn="conn" v-if="route.params.user !== undefined" />
  <ChatNotifictaion />
</template>

<style scoped>

</style>