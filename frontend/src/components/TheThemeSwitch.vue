<template>
  <label for="switch-theme" class="switch icon">
    <input
      type="checkbox"
      id="switch-theme"
      @click="toggleTheme"
      :checked="mode == 'dark'"
    />
    <span>
      <i>dark_mode</i>
    </span>
  </label>
</template>

<script setup lang="ts">
import { ref } from "vue";

const mode = ref(localStorage.getItem("mode") ?? "light");

function toggleTheme(): void {
  let newMode = mode.value == "light" ? "dark" : "light";

  mode.value = newMode;

  ui("mode", newMode);
  localStorage.setItem("mode", newMode);
}
</script>
