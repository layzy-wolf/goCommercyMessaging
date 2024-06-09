<script setup>

</script>

<template>
  <div class="video-wrapper">
    <video playsinline autoplay muted loop poster="/Particles.jpg">
      <source src="/Particles.mp4">
      You're browser don't support videos
    </video>
  </div>
</template>

<style scoped>
.video-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

video {
  object-fit: cover;
  height: 100%;
  width: 100%;
}
</style>