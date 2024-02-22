<script setup lang="ts">
defineProps({
  loading: { type: Boolean, default: false },
  icon: String,
  message: String,
});
</script>

<template>
  <article class="medium-height middle-align center-align">
    <div class="small-opacity">
      <progress class="loader circle surface" v-if="loading"></progress>
      <i class="extra" v-else>{{ icon }}</i>
      <h6>{{ message }}</h6>
    </div>
  </article>
</template>
