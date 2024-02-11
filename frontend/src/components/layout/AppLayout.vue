<script setup lang="ts">
import TheNavigation from './navigation/TheNavigation.vue';
import TheFooter from '../TheFooter.vue';
import TheContact from '../TheContact.vue';
</script>

<template>
  <TheNavigation />
  <div class="space"></div>
 
  <slot></slot>

  <div class="space"></div>
  <div class="divider"></div>
  <div class="space"></div>
  <TheContact />
  <div class="small-height"></div>
  <TheFooter />
</template>

<style scoped>
.space {
  margin-top: 3.5rem;
}
</style>
