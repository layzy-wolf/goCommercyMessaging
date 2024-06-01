<script setup>
</script>

<template>
  <div class="cont border-start">
    <div class="d-flex flex-column align-items-stretch flex-shrink-0">
      <div class="list-group list-group-flush border-bottom custom-scroll">
        <router-link to="/chat/user" class="list-group-item list-group-item-action active py-3 lh-sm" aria-current="true">
          <div class="d-flex w-100 align-items-center justify-content-between h-100">
            <strong class="mb-1">List group item heading</strong>
            <small>Wed</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>
        <router-link to="/chat/user" class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">List group item heading</strong>
            <small class="text-body-secondary">Tues</small>
          </div>
          <div class="col-10 mb-1 small">Some placeholder content in a paragraph below the heading and date.</div>
        </router-link>

      </div>
    </div>
  </div>

</template>

<style scoped>
  .cont {
    box-sizing: border-box;
    height: 100vh;
    width: 25vw;
    overflow: hidden;
  }
  .custom-scroll {
    overflow-y: scroll;
    max-height: 100vh !important;
    scrollbar-width: thin;
    scrollbar-color: rgba(0, 0, 0, 1) rgba(0, 0, 0, 0);
    transition: scrollbar-color .3s ease;
    -webkit-overflow-scrolling: touch;
    pointer-events: auto;
  }

  .list-group-item {
    background-color: rgba(0, 0, 0, 0);
  }

  .list-group-item.active {
    background-color: rgba(0, 0, 0, .5);
    border-color: rgba(0, 0, 0, 0);
  }
</style>