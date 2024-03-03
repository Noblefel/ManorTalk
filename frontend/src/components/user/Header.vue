<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { useUserStore } from "@/stores/user";
import type { RequestResponse } from "@/utils/api";
import { getAvatar } from "@/utils/helper";
import type { PropType } from "vue";

const props = defineProps({
  rr: { type: Object as PropType<RequestResponse>, required: true },
});

const authStore = useAuthStore();
const userStore = useUserStore();
</script>

<template>
  <div id="profile" v-if="userStore.viewedUser">
    <header>
      <img :src="getAvatar(userStore.viewedUser)" alt="Profile Avatar" />
      <h5 class="font-size-title">{{ userStore.viewedUser.name }}</h5>
      <h6 class="font-size-1 medium-opacity">
        # {{ userStore.viewedUser.username }}
      </h6>
    </header>

    <div class="row scroll no-space" v-if="!props.rr.errors">
      <RouterLink :to="{ name: 'profile', params: $route.params }">
        <button
          :class="$route.name == 'profile' ? 'secondary' : 'inverted'"
          :disabled="props.rr.loading"
        >
          <i>account_circle</i>
          Profile
        </button>
      </RouterLink>
      <RouterLink :to="{ name: 'profile.posts', params: $route.params }">
        <button
          :class="$route.name == 'profile.posts' ? 'secondary' : 'inverted'"
          :disabled="props.rr.loading"
        >
          <i>article</i>
          Posts
        </button>
      </RouterLink>
      <RouterLink :to="{ name: 'profile.edit', params: $route.params }">
        <button
          class="inverted"
          :disabled="props.rr.loading"
          v-if="authStore.authUser?.username == $route.params.username"
        >
          <i>edit</i>
          Edit
        </button>
      </RouterLink>
      <RouterLink :to="{ name: 'home' }">
        <button
          class="inverted"
          :disabled="props.rr.loading"
          v-if="authStore.authUser?.username == $route.params.username"
        >
          <i>settings</i>
          Settings
        </button>
      </RouterLink>
    </div>

    <slot :loading="props.rr.loading"></slot>
  </div>
  <div id="profile" v-else-if="props.rr.errors">
    <header class="small-opacity">
      <i class="font-size-2 responsive">warning</i>
      <h6>
        {{
          props.rr.status == 404
            ? "User not found"
            : "Unable to get user profile"
        }}
      </h6>
    </header>
  </div>
  <div id="profile" v-else>
    <header>
      <progress class="circle large tertiary-text"></progress>
    </header>
  </div>
</template>

<style scoped>
header {
  min-height: clamp(13rem, 11rem + 11vw, 20rem);
  border-radius: 12px;
  display: flex;
  justify-content: center;
  align-items: center;

  img {
    object-fit: cover;
    border-radius: 50%;
    width: 5.5rem;
    height: 5.5rem;
  }
}

#profile {
  padding: 1rem;
  max-width: 1050px;
  margin: auto;
}
</style>
