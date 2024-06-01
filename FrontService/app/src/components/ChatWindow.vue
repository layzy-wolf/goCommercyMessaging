<script setup>

import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import Message from "./Message.vue";

let messages = [
  {
    type: "received",
    user: "user",
    content: "hello world",
    timestamp: "13:39 24.04.24"
  },
  {
    type: "sent",
    user: "user2",
    content: "hello world",
    timestamp: "13:39 24.04.24"
  }
]
</script>

<template>
  <div class="cont">
    <div class="cont-wrapper custom-scroll border-end border-bottom d-flex flex-column justify-content-end">

        <Message v-for="message in messages" :message="message" />

    </div>
    <div class="message-zone text-muted d-flex justify-content-start align-items-center p-3 bg-dark">
      <input type="text" class="form-control form-control-lg" id="exampleFormControlInput1" placeholder="Type message">
      <a class="ms-3" href="#!"><FontAwesomeIcon :icon="['fas', 'paper-plane']" /></a>
    </div>
  </div>
</template>

<style scoped>
.cont {
  padding: 0;
  box-sizing: border-box;
  height: 100vh;
  width: 60vw;
  overflow: hidden;
  z-index: 3;
  margin: 0 10vw 0 10vw;
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
</style>