<script setup lang="ts">
import { RouterLink } from "vue-router";
import Profile from "./Profile.vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();
</script>

<template>
  <menu id="nav-menu" class="left">
    <Profile />

    <RouterLink :to="{name: 'blog.create'}"  v-if="authStore.isAuth">
      <button class="responsive no-margin secondary">
        Create Post
        <i>edit</i>
      </button>
    </RouterLink>

    <div class="space"></div>
    <div class="divider"></div>
    <div class="space"></div>

    <div v-if="!authStore.isAuth">
      <RouterLink
        :to="{ name: 'register' }"
        class="button responsive no-margin inverted"
      >
        Register
        <i>person_add</i>
      </RouterLink>
      <div class="space"></div>
      <RouterLink
        :to="{ name: 'login' }"
        class="button responsive no-margin primary"
      >
        Login
        <i>login</i>
      </RouterLink>
    </div>
    <div v-else>
      <RouterLink
        :to="{ name: 'logout' }"
        class="button responsive no-margin inverted"
      >
        Logout
        <i>Logout</i>
      </RouterLink>
    </div>
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
