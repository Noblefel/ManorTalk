<script setup lang="ts">
import TheNavigation from './navigation/TheNavigation.vue';  
</script>

<template>
  <TheNavigation />
  <div class="space"></div>
  <slot></slot>
  <div class="large-space"></div>
</template>

<style scoped>
.space {
  margin-top: 3.5rem;
}
</style>
