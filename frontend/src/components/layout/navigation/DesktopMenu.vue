<script setup lang="ts">
import { RouterLink } from "vue-router";
import Profile from "./Profile.vue";
</script>

<template>
  <menu id="nav-menu" class="left">
    <Profile />

    <button class="responsive no-margin secondary">
      Create Post
      <i>edit</i>
    </button>

    <div class="space"></div>
    <div class="divider"></div>
    <div class="space"></div>

    <RouterLink :to="{ name: 'register' }" >
      <button class="responsive no-margin inverted">
        Register
        <i>person_add</i>
      </button>
    </RouterLink>
    <div class="space"></div>
    <RouterLink :to="{ name: 'login' }" >
      <button class="responsive no-margin">
        Login
        <i>login</i>
      </button>
    </RouterLink>
  </menu>
</template>

<style scoped>
#nav-menu {
  bottom: -0.5rem;
  width: 19rem;
  background-color: var(--background);
  z-index: 2021 !important;
  border-radius: 10px;
  border: 1px solid color-mix(in srgb, var(--on-surface), transparent 75%);
  padding: 1rem;
}
a {
  padding: 0;

  &:hover {
    background-color: unset;
  }
}
</style>
