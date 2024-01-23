<script setup lang="ts">
import { RouterLink } from "vue-router";
import TheThemeSwitch from "./TheThemeSwitch.vue";
</script>

<template>
  <nav>
    <RouterLink to="/" class="row">
      <img class="logo" src="@/assets/images/logo_150.png" alt="App Logo" />
      <span class="font-700 font-size-1-5">ManorTalk.</span>
    </RouterLink>

    <div class="row l">
      <a class="item active"> Home </a>
      <a class="item"> Blog </a>
      <a class="item"> Discussion </a>
      <a class="item"> Categories</a>
    </div>

    <div class="row">
      <TheThemeSwitch />
      <button class="inverted">Login <i>login</i></button>
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

  .logo {
    height: 65px;
  }

  .item {
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
}
</style>
