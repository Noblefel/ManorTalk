<script setup lang="ts">
import { RouterLink } from "vue-router";
import { beerUi, activeRoute } from "@/utils/helper";
import MobileMenu from "./MobileMenu.vue";
import DesktopMenu from "./DesktopMenu.vue";
import TheThemeSwitch from "@/components/TheThemeSwitch.vue";
</script>

<template>
  <nav>
    <RouterLink to="/" class="row no-padding">
      <img class="logo" src="@/assets/images/logo_150.png" alt="App Logo" />
      <span class="font-700 font-size-1-5">ManorTalk</span>
    </RouterLink>

    <div class="row l">
      <RouterLink :to="{ name: 'home' }" :class="activeRoute('home')">
        Home
      </RouterLink>
      <RouterLink :to="{ name: 'home' }" :class="activeRoute('blog')">
        Blog
      </RouterLink>
      <RouterLink :to="{ name: 'home' }" :class="activeRoute('discussion')">
        Discussion
      </RouterLink>
      <RouterLink :to="{ name: 'home' }" :class="activeRoute('categories')">
        Categories
      </RouterLink>
    </div>

    <div class="row right-align">
      <TheThemeSwitch />

      <!-- DESKTOP -->
      <div class="l">
        <button class="menu-button background">
          <img
            src="@/assets/images/default_pfp_240.png"
            class="circle medium"
            alt=""
          />
        </button>
        <DesktopMenu />
      </div>

      <!-- MOBILE -->
      <div class="m s">
        <button class="menu-button" @click="beerUi('nav-mobile-menu')">
          <i>menu</i>
        </button>
        <MobileMenu />
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav {
  height: 4.5rem;
  padding: 1rem;
  border-bottom: 1px solid rgba(136, 136, 136, 0.247);
  position: fixed;
  top: 0;
  z-index: 5;
  background-color: var(--surface);
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.logo {
  height: 65px;
}

a {
  font-weight: 600;
  border-radius: 8px;
  padding: 0.5rem 1rem;

  &:hover {
    color: var(--on-secondary);
  }

  &.active {
    background-color: var(--secondary);
    color: var(--on-secondary);
  }
}

.menu-button {
  border-radius: 50%;
  margin: 0;
  padding: 0;
  width: 3rem;
  height: 3rem;

  i {
    width: 2rem;
    font-size: 1.75rem;
  }
}
</style>
