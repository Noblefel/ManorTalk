<script setup lang="ts">
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { usePostStore } from "@/stores/post";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import Filters from "@/components/blog/Filters.vue";
import PostList from "@/components/blog/PostList.vue";

const postStore = usePostStore();
const rr = ref(new RequestResponse());
const route = useRoute();

onMounted(() => {
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
      <Filters :loading="rr.loading"/>
      <div v-if="rr.errors">
        <article class="medium-height middle-align center-align">
          <div class="small-opacity">
            <i>warning</i>
            <h6>Unable to get posts</h6>
          </div>
        </article>
      </div>
      <div v-else-if="rr.loading">
        <div class="space"></div>
        <article class="medium-height middle-align center-align">
          <div class="small-opacity">
            <progress class="loader circle surface"></progress>
            <h6>Please wait...</h6>
          </div>
        </article>
      </div>
      <div class="s12 m12 l12" v-else-if="(rr.data as any)?.posts.length == 0">
        <article class="medium-height middle-align center-align">
          <div class="small-opacity">
            <i>error</i>
            <h6>Cannot find any post</h6>
          </div>
        </article>
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
