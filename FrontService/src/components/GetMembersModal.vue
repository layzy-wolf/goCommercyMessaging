<script setup>

import axios from "axios";
import {inject, onBeforeMount, ref} from "vue";
import {useRoute} from "vue-router";

const apiConfig = inject(["apiConfig"])
const openModalGetMembers = inject("openModalGetMembers")
const list = ref()
const route = useRoute()

function getMembers() {
  axios.post(apiConfig.http + "/group/get", {
    name: route.params.user
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  })
      .then(res => {
        if (res.data.members !== undefined) {
          list.value = res.data.members
        }
      })
}

onBeforeMount(() => {
  getMembers()
})

function removeUser(user) {
  axios.post(apiConfig.http + "/group/remove", {
    group: route.params.user,
    remove: user
  }, {
    headers: {
      Authorization: localStorage.getItem("token")
    }
  })
      .then(res => {
        if (res.data.success) {
          getMembers()
        }
      })
}

function close() {
  openModalGetMembers.value = false
}

</script>

<template>
  <div class="modal bg-black bg-opacity-25" style="display: block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title text-light">Члены Группы</h5>
          <button type="button" class="btn-close" @click.prevent="close"></button>
        </div>
        <div class="modal-body">
          <div class="d-flex flex-column w-100 gap-2 overflow-auto overflow-x-hidden custom-scroll"
               style="max-height: 200px">
            <div v-for="user in list" class="list-group-item py-3 lh-sm"
                 aria-current="true">
              <div class="d-flex w-100 align-items-center justify-content-between h-100 px-3">
                <strong class="mb-1 text-light">{{ user }}</strong>
                <button class="btn btn-outline-danger rounded-0" @click.prevent="removeUser(user)">Удалить</button>
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