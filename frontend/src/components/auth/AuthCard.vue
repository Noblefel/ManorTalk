<script setup lang="ts">
import TheThemeSwitch from "./../TheThemeSwitch.vue";
import { RouterLink } from "vue-router";

defineProps({
  title: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <article>
    <div>
      <RouterLink :to="{ name: 'home' }" class="home-button">
        <i>home</i>
      </RouterLink>
      <img src="@/assets/images/logo_150.png" alt="App icon" class="center" />
      <TheThemeSwitch class="absolute top right" />
    </div>

    <div class="divider"></div>
    <div></div>
    <h2 class="font-size-title center-align">
      {{ title }}
    </h2>

    <slot></slot>
  </article>
</template>

<style scoped>
article {
  max-width: 95%;
  width: 27.5rem;
  padding-bottom: 2rem;
  top: -1rem;
}

.home-button {
  padding: 0.5rem;
  position: absolute;
  background-color: var(--secondary);
  color: var(--on-secondary);
}

img {
  width: 5rem;
  opacity: 0.9;
}
</style>
