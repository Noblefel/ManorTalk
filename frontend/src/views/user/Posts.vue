<script setup lang="ts">
import Header from "@/components/user/Header.vue";
import PostCard from "@/components/PostCard.vue";
import ResponseCard from "@/components/ResponseCard.vue";
import { usePostStore, type Post } from "@/stores/post";
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { useUserStore } from "@/stores/user";
import { onBeforeRouteUpdate, useRoute, RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const route = useRoute();
const rr = ref(new RequestResponse());
const rrPosts = ref(new RequestResponse());
const posts = ref<Post[]>([]);
const cursor = ref(1);
const postStore = usePostStore();
const userStore = useUserStore();
const authStore = useAuthStore();

onMounted(() => {
  if (userStore.viewedUser) {
    load();
    return;
  }

  userStore.fetchProfile(route, rr.value)?.then(() => {
    load();
  });
});

onBeforeRouteUpdate((to) => {
  userStore.fetchProfile(to, rr.value);
});

function load() {
  postStore.fetchProfilePosts(rrPosts.value, posts, cursor);
}
</script>

<template>
  <Header :rr="rr">
    <div class="space"></div>
    <div class="center-align margin">
      <RouterLink
        :to="{ name: 'blog.create' }"
        class="button large"
        v-if="authStore.authUser?.username == $route.params.username"
      >
        Create Posts&nbsp; ✏️
      </RouterLink>
    </div>
    <div v-if="posts && posts.length">
      <div class="grid">
        <div class="s12 m6 l4" v-for="post in posts">
          <PostCard
            :post="post"
            :with-excerpt="true"
            :separate-image="true"
            image-height="10rem"
            background-color="var(--background)"
            image="/src/assets/images/stock_1.jpg"
          />
          <div class="space"></div>
        </div>
      </div>
      <div class="center-align">
        <div class="space"></div>
        <button @click="load" v-if="!rrPosts.loading">Load More</button>
        <progress v-else class="circle"></progress>
      </div>
    </div>
    <div v-else-if="rrPosts.loading">
      <ResponseCard :loading="true" message="Please wait..." />
    </div>
    <div v-else-if="rrPosts.errors">
      <ResponseCard icon="warning" message="Unable to get posts" />
    </div>
    <div v-else>
      <ResponseCard icon="article" message="Empty" />
    </div>
  </Header>
</template>
