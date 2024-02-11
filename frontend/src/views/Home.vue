<script setup lang="ts">
import Hero from "@/components/home/Hero.vue";
import Topics from "@/components/home/Topics.vue";
import NewestPosts from "@/components/home/NewestPosts.vue";
import PostCard from "@/components/PostCard.vue";
</script>

<template>
  <div>
    <Hero />
    <Topics />
    <div class="divider"></div>

    <section id="home-featured">
      <div class="wrapper">
        <h4 class="font-size-title wrap">
          <span class="orange-text">Featured</span>
          Posts This Week
        </h4>
        <div class="grid">
          <PostCard
            class="s12 m6 l6"
            image="/src/assets/images/stock_1.jpg"
            :with-border="true"
            :with-excerpt="true"
            :with-author="true"
            :with-stats="true"
          />
          <PostCard
            class="s12 m6 l6"
            :with-border="true"
            :with-excerpt="true"
            :with-author="true"
            :with-stats="true"
          />
        </div>
      </div>
    </section>

    <div class="large-space"></div>

    <NewestPosts />

    <div class="large-space"></div>

    <section id="home-discussions">
      <div class="wrapper">
        <h4 class="font-size-title wrap">
          <span class="orange-text">Trending</span>
          Discussions
        </h4>
        <div class="grid">
          <article class="glow s12 m6 l6 no-margin">
            <h6>Discussion #1</h6>
            <p class="small-opacity">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
              sapien luctus libero ultricies egestas vel vitae ligula. Aliquam
              posuere condimentum ex, quis luctus libero interdum eget
            </p>
          </article>
          <article class="glow s12 m6 l6 no-margin">
            <h6>Discussion #2</h6>
            <p class="small-opacity">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
              sapien luctus libero ultricies egestas vel vitae ligula. Aliquam
              posuere condimentum ex, quis luctus libero interdum eget
            </p>
          </article>
          <article class="glow s12 m6 l6 no-margin">
            <h6>Discussion #3</h6>
            <p class="small-opacity">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
              sapien luctus libero ultricies egestas vel vitae ligula. Aliquam
              posuere condimentum ex, quis luctus libero interdum eget
            </p>
          </article>
          <article class="glow s12 m6 l6 no-margin">
            <h6>Discussion #4</h6>
            <p class="small-opacity">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
              sapien luctus libero ultricies egestas vel vitae ligula. Aliquam
              posuere condimentum ex, quis luctus libero interdum eget
            </p>
          </article>
        </div>
      </div>
      <div class="padding center-align">
        <button class="inverted">
          View More
          <i class="font-size-2">arrow_right</i>
        </button>
      </div>
    </section> 
  </div>
</template>
