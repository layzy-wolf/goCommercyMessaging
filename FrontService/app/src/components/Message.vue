<script setup>
  const message = defineProps(["message"])
</script>

<template>
  <div class="d-grid message rounded text-light my-2 mx-2 p-2 w-100" :class="message.message.type">
    <span class="message-header badge text-wrap text-dark rounded-pill bg-light">
      @{{message.message.user}}
    </span>
    <span class="message-time badge text-dark opacity-75">
      {{message.message.timestamp}}
    </span>
    <h5 class="message-content fs-6 my-2 text-wrap text-break">
      {{message.message.content}}
    </h5>
    <svg width="9" height="20" class="svg-appendix">
      <g fill="#495057" fill-rule="evenodd">
      <path d="M6 17H0V0c.193 2.84.876 5.767 2.05 8.782.904 2.325 2.446 4.485 4.625 6.48A1 1 0 016 17z" filter="url(#messageAppendix)"></path>
      <path d="M6 17H0V0c.193 2.84.876 5.767 2.05 8.782.904 2.325 2.446 4.485 4.625 6.48A1 1 0 016 17z" class="corner"></path>
    </g>
    </svg>

  </div>

</template>

<style scoped>
  .message {
    width: fit-content !important;
    max-width: 25vw;
    min-width: 5vw;
    background: #495057;
    grid-template-columns: auto min-content;
    position: relative;
  }

  .message-header {
    grid-area: a;
  }

  .message-time {
    grid-area: b;
  }

  .message-content {
    grid-area: c;
  }

  .received {
    grid-template-areas:  "a b"
                          "c c";
    margin-right: auto;
  }

  .received > .message-header {
    margin-right: auto;
  }

  .received > .message-content {
    margin-right: auto;
    margin-left: 10px;
    grid-column: 1/3;
  }

  .received > .svg-appendix {
    position: absolute;
    bottom: -5px;
    left: -8px;
    rotate: 270deg;
  }

  .sent {
    margin-left: auto !important;
    grid-template-areas:  "b a"
                          "c c";
  }

  .sent > .message-header {
    margin-left: auto;
  }

  .sent > .message-time {
    margin-right: auto;
  }

  .sent > .message-content {
    margin-left: auto;
    margin-right: 10px;
    grid-column: 1/3;
  }

  .sent > .svg-appendix {
    position: absolute;
    bottom: -5px;
    right: -8px;
    transform: scale(-1, 1);
    rotate: 90deg;
  }
</style>