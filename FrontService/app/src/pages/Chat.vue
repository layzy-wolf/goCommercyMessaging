<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import LeftBar from "../components/LeftBar.vue";
import {computed, watch} from "vue";
import {router} from "../router.js";
import {useRoute} from "vue-router";
import ChatBar from "../components/ChatBar.vue";
import ChatWindow from "../components/ChatWindow.vue";
const route = useRoute()

watch(() => route.params.user, (newU, oldU) => {
  route.params.user = newU
})
</script>

<template>
  <LeftBar />
  <ChatBar />
  <ChatWindow v-if="route.params.user !== undefined" />
</template>

<style scoped>

</style>