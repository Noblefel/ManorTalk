<script setup lang="ts">
import PostCard from "../PostCard.vue";
import { usePostStore } from "@/stores/post";

const postStore = usePostStore();

const { data, errors, status } = postStore.fetchLatestPosts();
</script>

<template>
  <section id="home-new">
    <div class="wrapper"> 
      <h4 class="font-size-title wrap">
        <span class="orange-text">Newest</span>
        Posts
      </h4>
      <div v-if="errors">
        <div class="space"></div>
        <article class="small-height middle-align center-align">
          <div class="small-opacity" v-if="status == 404">
            <i>error</i>
            <h6>Empty</h6>
          </div>
          <div class="small-opacity" v-else>
            <i>warning</i>
            <h6>Unable to get newest posts</h6>
          </div>
        </article>
      </div>
      <div class="row scroll" v-else-if="data">
        <PostCard
        v-for="post in data.posts"
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
      <div v-else>
        <div class="space"></div>
        <article class="small-height middle-align center-align">
          <div class="small-opacity">
            <progress class="loader circle surface"></progress>
            <h6>Please wait...</h6>
          </div>
        </article>
      </div>

      <button class="inverted absolute top right">
        View More
        <i class="font-size-2">arrow_right</i>
      </button>
    </div>
  </section>
</template>
