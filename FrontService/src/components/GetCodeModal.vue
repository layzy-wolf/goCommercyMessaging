<script setup>

import axios from "axios";
import {inject, onBeforeMount, ref} from "vue";
import {useRoute} from "vue-router";

const apiConfig = inject("apiConfig")
const openModalGetCode = inject("openModalGetCode")
const route = useRoute()
let code = ref()

onBeforeMount(() => {
  axios.post(apiConfig.http + "/group/code", {
    name: route.params.user
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  })
      .then(res => {
        if (res.data.error !== null && res.data.error !== "") {
          code.value = res.data.code
        } else {
          code.value = res.data.error
        }
      })
})

function close() {
  openModalGetCode.value = false
}
</script>

<template>
  <div class="modal bg-black bg-opacity-25" style="display: block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title text-light">Код Группы</h5>
          <button type="button" class="btn-close" @click.prevent="close"></button>
        </div>
        <div class="modal-body">
          <span class="d-block text-light overflow-y-hidden overflow-x-scroll custom-scroll w-100">{{code}}</span>
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