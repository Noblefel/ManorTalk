<script setup lang="ts">
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { usePostStore } from "@/stores/post";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import Filters from "@/components/blog/Filters.vue";
import PostList from "@/components/blog/PostList.vue";
import ResponseCard from "@/components/ResponseCard.vue";

const postStore = usePostStore();
const rr = ref(new RequestResponse());
const route = useRoute();

onMounted(() => {
  window.scrollTo(0, 0);
  postStore.fetchPosts(rr.value, route.fullPath);
});

onBeforeRouteUpdate((to) => {
  postStore.fetchPosts(rr.value, to.fullPath);
});
</script>

<template>
  <div id="blog">
    <header>
      <h1>Blog 📝</h1>
    </header>
    <div class="wrapper">
      <Filters :loading="rr.loading" />
      <div v-if="rr.errors">
        <ResponseCard icon="warning" message="Unable to get posts" />
      </div>
      <div v-else-if="rr.loading">
        <ResponseCard :loading="true" message="Please wait..." />
      </div>
      <div class="s12 m12 l12" v-else-if="(rr.data as any)?.posts.length == 0">
        <ResponseCard icon="error" message="Cannot find any posts" />
      </div>
      <PostList
        v-else
        :posts="(rr.data as any)?.posts"
        :pagination_meta="(rr.data as any)?.pagination_meta"
      />
    </div>
  </div>
</template>

<style scoped>
header {
  height: 10rem;
  display: flex;
  align-items: center;
}

.wrapper {
  max-width: 950px;
  padding: 0.5rem;
}
</style>
