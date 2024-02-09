<script setup lang="ts">
import Pagination from "@/components/Pagination.vue";
import PostCard from "@/components/PostCard.vue";
import type { PaginationMeta, Post } from "@/stores/post";
import type { PropType } from "vue";

defineProps({
  posts: {
    type: Array as PropType<Post[]>,
  },
  pagination_meta: {
    type: Object as PropType<PaginationMeta>,
  },
});
</script>

<template>
  <div class="grid">
    <div class="s12 m12 l12" v-if="pagination_meta">
      <h6>
        <span class="font-size-1-5">🔎&nbsp;</span>
        Found {{ Intl.NumberFormat().format(pagination_meta.total) }} posts
      </h6>
    </div>
    <div class="s12 m6 l6" v-for="post in posts">
      <PostCard
        class="s6 m4 l3"
        :post="post"
        :with-author="true"
        :with-excerpt="true"
        :separate-image="true"
        image-height="clamp(10rem, calc(5rem + 15vw), 18rem)"
        background-color="var(--background)"
        image="/src/assets/images/stock_1.jpg"
      />
    </div>
    <div class="s12 m12 l12" v-if="pagination_meta">
      <div class="space"></div>
      <Pagination :pagination-meta="pagination_meta" />
    </div>
  </div>
</template>

<style scoped>
h6 {
  padding: 0 1rem;
  font-size: 1rem;
}
</style>
