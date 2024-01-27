<template>
  <div class="auth-background">
    <slot></slot>
  </div>
</template>

<style scoped>
.auth-background {
  display: flex;
  justify-content: center;
  align-items: center; 
  min-height: 100vh;
  padding: 2rem 0;
}
</style>
