<script setup lang="ts">
import { onBeforeMount } from "vue";
import { RouterView } from "vue-router"; 

onBeforeMount(() => {
  ui("mode", localStorage.getItem("mode") ?? "light");
});
</script>

<template>
  <component :is="$route.meta.layout || 'div'">
    <RouterView v-slot="slotProps">
      <Transition name="route" mode="out-in">
        <component :is="slotProps.Component"></component>
      </Transition>
    </RouterView>
  </component>
</template>

<style>
.route-enter-active {
  animation: fading 0.5s ease
}

.route-leave-from {
  opacity: 1;
}

.route-leave-to {
  opacity: 0;
}

.route-leave-active {
  transition: opacity 0.5s ease;
}

@keyframes fading {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

</style>