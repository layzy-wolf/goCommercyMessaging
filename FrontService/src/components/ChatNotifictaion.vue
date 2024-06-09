<script setup>


import {inject, onMounted, ref, watch} from "vue";

const message = inject("message")
const modalOpen = inject("modal")
let modal
onMounted(() => {
  modal = document.querySelector("#modal")
})

watch(modalOpen, () => {
  modal.classList.add("show")
  setTimeout(() => {
    modal.classList.remove("show")
  }, 2000)
})

</script>

<template>
<div class="modalCustom fade text-light rounded bg-black p-2" id="modal">
  <span class="message-header badge text-wrap text-dark rounded-pill bg-light" v-if="message !== undefined">
      {{message.from}}
  </span>
  <span class="message-time badge text-dark">
      Новое
    </span>
  <h5 class="message-content fs-6 my-2 text-wrap text-break" v-if="message !== undefined">
    {{message.message}}
  </h5>
</div>
</template>

<style scoped>
  .modalCustom {
    position: absolute;
    padding-right: 1.25rem;
    bottom: 5%;
    right: 5%;
    z-index: 20;
  }

  .show {
    opacity: 1 !important;
  }
</style>