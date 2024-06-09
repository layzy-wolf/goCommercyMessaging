<script setup>
import {inject} from "vue";
import axios from "axios";

const apiConfig = inject("apiConfig")
const openModalEnterGroup = inject("openModalEnterGroup")
const updateChat = inject("updateChat")

let token

function enter() {
  if (token !== undefined && token !== "") {
    axios.post(apiConfig.http + "/group/add", {
      token: token
    }, {
      headers: {
        Authorization: localStorage.getItem("token")
      }
    }).then(res => {
      if (res.data.success) {
        updateChat.value = !updateChat.value
        close()
      }
    }).catch()
  }
}

function close() {
  openModalEnterGroup.value = false
}
</script>

<template>
  <div class="modal bg-black bg-opacity-25" style="display: block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title text-light">Войти в Группу</h5>
          <button type="button" class="btn-close" @click.prevent="close"></button>
        </div>
        <div class="modal-body">

          <form action="#">
            <div class="input-group">

              <div class="form-floating">
                <input type="text" class="form-control" placeholder="Group Token" v-model="token">
                <label class="text-light" for="search">Group Token</label>
              </div>
              <button class="btn btn-outline-secondary" type="submit" @click.prevent="enter">Войти</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>