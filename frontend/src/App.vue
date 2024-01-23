<script setup lang="ts">
import { onBeforeMount } from "vue";
import { RouterView } from "vue-router";
import Home from "@/views/Home.vue";
import TheNavigation from "./components/TheNavigation.vue";
import TheFooter from "./components/TheFooter.vue";

onBeforeMount(() => {
  ui("mode", localStorage.getItem("mode") ?? "light");
});
</script>

<template>
  <TheNavigation />
  <div class="space"></div>
  <Home />

  <div class="small-height"></div>
  <TheFooter />
</template>

<style scoped>
.space {
  margin-top: 3.5rem;
}
</style>
