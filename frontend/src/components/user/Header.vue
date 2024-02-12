<script setup lang="ts">
import type { PropType } from "vue";
import type { User } from "./../../stores/user";
import { getAvatar } from "@/utils/helper";
import { useAuthStore } from "@/stores/auth";
import { auth } from "@/router/middleware";

const authStore = useAuthStore();

defineProps({
  user: {
    type: Object as PropType<User>,
    required: true,
  },
  errors: Boolean,
  loading: Boolean,
  status: Number,
});
</script>

<template>
  <div id="profile" v-if="user">
    <header>
      <img :src="getAvatar(user)" alt="Profile Avatar" />
      <h5 class="font-size-title">{{ user.name }}</h5>
      <h6 class="font-size-1 medium-opacity"># {{ user.username }}</h6>
    </header>

    <div class="row scroll" v-if="!errors">
      <button class="secondary" :disabled="loading">
        <i>account_circle</i>
        Profile
      </button>
      <button class="inverted" :disabled="loading">
        <i>article</i>
        Posts
      </button>
      <button
        class="inverted"
        :disabled="loading"
        v-if="authStore.authUser?.username == $route.params.username"
      >
        <i>edit</i>
        Edit
      </button>
      <button
        class="inverted"
        :disabled="loading"
        v-if="authStore.authUser?.username == $route.params.username"
      >
        <i>settings</i>
        Settings
      </button>
    </div>

    <slot></slot>
  </div>
  <div id="profile" v-else-if="errors">
    <header class="small-opacity">
      <i class="font-size-2 responsive">warning</i>
      <h6>
        {{ status == 404 ? "User not found" : "Unable to get user profile" }}
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
    border-radius: 50%;
    width: 5rem;
  }
}

#profile {
  padding: 1rem;
  max-width: 1050px;
  margin: auto;
}
</style>
