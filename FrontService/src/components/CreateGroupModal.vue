<script setup>

import axios from "axios";
import {inject} from "vue";

const apiConfig = inject("apiConfig")
const openModalCreateGroup = inject("openModalCreateGroup")
const updateChat = inject("updateChat")

let name

function create() {
  if (name !== undefined && name !== "") {
    axios.post(apiConfig.http + "/group/register", {
      name: name
    }, {
      headers: {
        Authorization: localStorage.getItem("token")
      }
    }).then(() => {
      updateChat.value = !updateChat.value
      close()
    })
  }
}

function close() {
  openModalCreateGroup.value = false
}
</script>

<template>
  <div class="modal bg-black bg-opacity-25" style="display: block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title text-light">Создать Группу</h5>
          <button type="button" class="btn-close" @click.prevent="close"></button>
        </div>
        <div class="modal-body">
          <form action="#">
            <div class="input-group">

              <div class="form-floating">
                <input type="text" class="form-control" placeholder="Имя Группы" v-model="name">
                <label class="text-light" for="search">Имя Группы</label>
              </div>
              <button class="btn btn-outline-secondary" type="submit" @click.prevent="create">Создать</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>