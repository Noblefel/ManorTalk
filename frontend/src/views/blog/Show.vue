<script setup lang="ts">
import ResponseCard from "@/components/ResponseCard.vue";
import PostCard from "@/components/PostCard.vue";
import Post from "@/components/blog/Post.vue";
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
import { usePostStore } from "@/stores/post";

const rr = ref(new RequestResponse());
const ps = usePostStore();
const route = useRoute();

onMounted(() => {
  window.scrollTo(0, 0);
  ps.fetchPost(rr, route);
});

onBeforeRouteUpdate((to) => {
  window.scrollTo(0, 0);
  ps.fetchPost(rr, to);
});
</script>

<template>
  <div class="wrapper">
    <main v-if="ps.viewedPost" id="post">
      <Post :post="ps.viewedPost" />

      <div class="large-space"></div>
      <div class="divider"></div>
      <div class="large-space"></div>

      <div class="grid">
        <article class="max glow s12 m6 l6 no-margin">
          <i>chevron_left</i>
          Previous post

          <PostCard :with-excerpt="true" :post="ps.viewedPost" />
        </article>
        <article class="max glow s12 m6 l6 no-margin">
          <div class="right-align">
            Next post
            <i>chevron_right</i>
          </div>

          <PostCard :with-excerpt="true" :post="ps.viewedPost" />
        </article>
      </div>
    </main>
    <div v-else-if="rr.errors">
      <ResponseCard
        icon="error"
        message="Cannot find that post"
        v-if="rr.status == 404"
      />
      <ResponseCard
        icon="warning"
        message="We're unable to get that post"
        v-else
      />
    </div>
    <div v-else>
      <ResponseCard :loading="true" message="Please wait..." />
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  max-width: 950px;
  padding: 1rem 0.5rem;
}
</style>
