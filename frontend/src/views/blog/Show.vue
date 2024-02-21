<script setup lang="ts">
import ResponseCard from "@/components/ResponseCard.vue";
import PostCard from "@/components/PostCard.vue";
import Post from "@/components/blog/Post.vue";
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { onBeforeRouteUpdate, useRoute } from "vue-router";

const rr = ref(new RequestResponse());
const route = useRoute();

onMounted(() => {
  rr.value.useApi("GET", "/posts/" + route.params.slug);
});

onBeforeRouteUpdate((to) => {
  rr.value.useApi("GET", "/posts/" + to.params.slug);
});
</script>

<template>
  <div class="wrapper">
    <main v-if="rr.data" id="post">
      <Post :post="(rr.data as any)" />

      <div class="large-space"></div>
      <div class="divider"></div>
      <div class="large-space"></div>

      <div class="grid">
        <article class="max glow s12 m6 l6 no-margin">
          <i>chevron_left</i>
          Previous post

          <PostCard :with-excerpt="true" :post="rr.data" />
        </article>
        <article class="max glow s12 m6 l6 no-margin">
          <div class="right-align">
            Next post
            <i>chevron_right</i>
          </div>

          <PostCard :with-excerpt="true" :post="rr.data" />
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
