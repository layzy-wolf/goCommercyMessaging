<script setup>

import axios from "axios";
import {inject, ref} from "vue";

const apiConfig = inject("apiConfig")
const openModalAddContact = inject("openModalAddContact")
const updateChat = inject("updateChat")
const list = ref()

function search(ev) {
  if (ev.target.value === "") {
    list.value = {}
    return
  }

  axios.post(apiConfig.http + "/search", {
    condition: ev.target.value
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  }).then(res => {
    if (res.data !== null) {
      list.value = res.data.users
    }
  })
}

function addContact(user) {
  axios.post(apiConfig.http + "/add/contact", {
    chat: user
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  }).then(() => {
    updateChat.value = !updateChat.value
    close()
  })
}

function close() {
  openModalAddContact.value = false
}

</script>

<template>
  <div class="modal bg-black bg-opacity-25" style="display: block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title text-light">Добавить Котакт</h5>
          <button type="button" class="btn-close" @click.prevent="close"></button>
        </div>
        <div class="modal-body">
          <div class="form-floating">
            <input type="text" class="form-control" id="search" placeholder="Поиск" @input="search">
            <label class="text-light" for="search">Поиск</label>
          </div>
        </div>
        <div class="modal-footer">
          <div class="d-flex flex-column w-100 gap-2 overflow-auto overflow-x-hidden custom-scroll"
               style="max-height: 200px">
            <div v-for="user in list" class="list-group-item py-3 lh-sm"
                 aria-current="true">
              <div class="d-flex w-100 align-items-center justify-content-between h-100 px-3">
                <strong class="mb-1 text-light">{{ user }}</strong>
                <button class="btn btn-outline-danger rounded-0" @click.prevent="addContact(user)">Добавить</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.custom-scroll {
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 1) rgba(0, 0, 0, 0);
  transition: scrollbar-color .3s ease;
  -webkit-overflow-scrolling: touch;
  pointer-events: auto;
}
</style>