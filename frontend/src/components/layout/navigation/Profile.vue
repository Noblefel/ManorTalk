<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { getAvatar } from "@/utils/helper";
import { RouterLink } from "vue-router";

const authStore = useAuthStore();
</script>

<template>
  <div class="nav-profile">
    <div class="row responsive">
      <img
        :src="getAvatar(authStore.authUser)"
        alt="Profile avatar"
        width="50px"
        height="50px"
        class="circle"
      />
      <div class="max right-align row">
        <RouterLink
          :to="{
            name: 'profile',
            params: { username: authStore.authUser?.username },
          }"
          class="button no-margin circle secondary"
          v-if="authStore.isAuth"
          data-ui="#nav-mobile-menu"
        >
          <i class="fill">manage_accounts</i>
          <div class="tooltip bottom">Profile</div>
        </RouterLink>
        <button class="no-margin circle m s" data-ui="#nav-mobile-menu">
          <i>close</i>
        </button>
      </div>
    </div>
    <div v-if="authStore.isAuth">
      <p class="no-margin" v-if="authStore.authUser?.name">
        {{ authStore.authUser.name }}
      </p>
      <p class="no-margin font-size-0-9 small-opacity">
        #{{ authStore.authUser?.username }}
      </p>
    </div>
    <div v-else>
      <p class="no-margin">Guest</p>
    </div>
  </div>
</template>

<style scoped>
.nav-profile {
  margin-bottom: 0.5rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  font-weight: 600;
}
</style>
