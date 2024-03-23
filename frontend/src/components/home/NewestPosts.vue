<script setup lang="ts">
import PostCard from "../PostCard.vue";
import ResponseCard from "../ResponseCard.vue";
import { RequestResponse } from "@/utils/api";
import { usePostStore } from "@/stores/post";
import { onMounted, ref } from "vue";

const rr = ref(new RequestResponse());
const ps = usePostStore();

onMounted(() => {
  ps.fetchHomePosts(rr)
});
</script>

<template>
  <section id="home-new">
    <div class="wrapper">
      <h4 class="font-size-title wrap">
        <span class="orange-text">Newest</span>
        Posts
      </h4>
      <div v-if="rr.errors">
        <div class="space"></div>
        <ResponseCard icon="error" message="Empty" v-if="rr.status == 404" />
        <ResponseCard icon="warning" message="Unable to get new posts" v-else />
      </div>
      <div v-else-if="rr.loading">
        <div class="space"></div>
        <ResponseCard :loading="true" message="Getting posts..." />
      </div>
      <div class="row scroll" v-else>
        <PostCard
          v-for="post in ps.latestPosts"
          class="s6 m4 l3"
          :key="post.id"
          :post="post"
          width="20rem"
          :with-excerpt="true"
          :separate-image="true"
          image-height="8rem"
          background-color="var(--surface)"
          image="/src/assets/images/stock_1.jpg"
        />
      </div> 

      <button class="inverted absolute top right">
        View More
        <i class="font-size-2">arrow_right</i>
      </button>
    </div>
  </section>
</template>
