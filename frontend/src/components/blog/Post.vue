<script setup lang="ts">
import { getAvatar } from "@/utils/helper";
import type { Post } from "@/stores/post";
import { type PropType } from "vue";
import Markdown from "../Markdown.vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();

defineProps({
  post: { type: Object as PropType<Post>, required: true },
});
</script>

<template>
  <header id="post-header">
    <h1>{{ post.title }}</h1>
    <div class="tags center-align">
      <p># {{ post.category.name }}</p>
    </div>
  </header>

  <div class="row wrap">
    <div class="author">
      <img
        :src="getAvatar(post.user)"
        :alt="`${post.user.username} profile avatar`"
      />
      <div>
        <RouterLink
          :to="{ name: 'profile', params: { username: post.user.username } }"
          class="font-700"
        >
          {{ post.user.name ?? post.user.username }}
        </RouterLink>
        <p class="small-text font-600" v-if="post.created_at">
          {{ new Date(post.created_at).toUTCString() }}
        </p>
      </div>
    </div>
    <div
      class="max right-align"
      v-if="authStore.authUser?.username == post.user?.username"
    >
      <RouterLink
        :to="{ name: 'blog.edit', params: $route.params }"
        class="button secondary"
      >
        <i>edit</i>
        Edit
      </RouterLink>
    </div>
  </div>

  <div class="content-wrapper small-padding">
    <!-- <img src="@/assets/images/stock_1.jpg" alt="" /> -->

    <p class="excerpt italic">
      {{ post.excerpt }}
    </p>

    <div class="space"></div>

    <Markdown :content="post.content" />
  </div>
</template>

<style scoped>
header {
  min-height: 12rem;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;

  h1 {
    font-size: clamp(1.5rem, calc(1rem + 3vw), 3rem);
  }

  .tags {
    display: flex;
    gap: 1rem;
    color: white;
    width: 100%;

    > * {
      font-size: 1rem;
      background-image: linear-gradient(to right, rgb(56, 53, 223), #8f24d6);
      padding: 0.3rem 0.5rem;
      border-radius: 5px;
    }
  }
}

.author {
  margin: 1rem 0.5rem;
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;

  img {
    width: 55px;
    height: 55px;
    object-fit: cover;
    border-radius: 50%;
  }
}

.content-wrapper {
  .excerpt {
    font-size: 1.1rem;
    font-weight: 500;
    opacity: 0.7;
  }

  img {
    width: 100%;
    height: clamp(200px, calc(100px + 40vw), 450px);
    object-fit: cover;
    margin: 1rem 0;
    border-radius: 15px;
  }
}
</style>
