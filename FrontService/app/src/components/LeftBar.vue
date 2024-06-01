<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import * as jwt from "jsonwebtoken-esm"
import {router} from "../router.js";
import {onBeforeMount} from "vue";

onBeforeMount(() => {
  if (localStorage.getItem("token") === null || localStorage.getItem("token") === "") {
    router.push("/sign")
  }
})

let user = jwt.decode(localStorage.getItem("token"))

function logout() {
  localStorage.removeItem("token")
  router.push("/sign")
}
</script>

<template>
  <div class="d-flex flex-column flex-shrink-0 justify-content-between" style="width: 4.5rem; height: 100%">
    <FontAwesomeIcon icon="fa-solid fa-comments-dollar" class="bi d-block p-3 link-body-emphasis text-decoration-none" width="40" height="32"  />
    <div class="dropdown border-top">
      <a href="#" class="d-flex align-items-center justify-content-center p-3 link-body-emphasis text-decoration-none dropdown-toggle" data-bs-toggle="dropdown" aria-expanded="false">
        <h5>@</h5>
      </a>
      <ul class="dropdown-menu text-small shadow" style="">
        <li class="dropdown-header border-bottom">
          <h5 class="text-light">@{{ user.login }}</h5>
        </li>
        <li><a class="dropdown-item" href="#" @click.prevent="logout">Sign out</a></li>
      </ul>
    </div>
  </div>
</template>

<style scoped>

</style>