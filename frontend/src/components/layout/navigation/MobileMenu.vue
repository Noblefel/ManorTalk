<script setup lang="ts">
import { activeRoute } from "@/utils/helper";
import Profile from "./Profile.vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();
</script>

<template>
  <dialog id="nav-mobile-menu" class="right">
    <Profile />

    <button class="responsive no-margin secondary">
      Create Post
      <i>edit</i>
    </button>

    <div class="space"></div>

    <div class="divider"></div>

    <div class="links">
      <RouterLink
        :to="{ name: 'home' }"
        :class="activeRoute('home')"
        data-ui="#nav-mobile-menu"
      >
        <i>cottage</i>
        Home
      </RouterLink>
      <RouterLink
        :to="{ name: 'blog' }"
        :class="activeRoute('blog')"
        data-ui="#nav-mobile-menu"
      >
        <i>newspaper</i>
        Blog
      </RouterLink>
      <RouterLink
        :to="{ name: 'home' }"
        :class="activeRoute('discussion')"
        data-ui="#nav-mobile-menu"
      >
        <i>forum</i>
        Discussion
      </RouterLink>
      <RouterLink
        :to="{ name: 'home' }"
        :class="activeRoute('categories')"
        data-ui="#nav-mobile-menu"
      >
        <i>category</i>
        Categories
      </RouterLink>
    </div>

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
      <RouterLink :to="{ name: 'login' }" class="button responsive no-margin">
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
        <i>logout</i>
      </RouterLink>
    </div>
  </dialog>
</template>

<style scoped>
#nav-mobile-menu {
  border-radius: 0;
  background-color: var(--background);
  max-width: 400px;

  @media screen and (max-width: 500px) {
    width: 100%;
    max-width: 100%;
  }
}

.links {
  display: flex;
  padding: 1rem 0;
  flex-direction: column;
  gap: 0.5rem;

  font-weight: 600;

  a {
    opacity: 0.8;
    justify-content: flex-start;
    padding: 0.75rem 1rem;
    border-radius: 8px;
    gap: 0.5rem;

    &:hover,
    &.active {
      opacity: 1;
      background-color: var(--surface);
    }
  }
}
</style>
