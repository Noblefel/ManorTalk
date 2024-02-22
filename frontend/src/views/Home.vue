<script setup lang="ts">
import Topics from "@/components/home/Topics.vue";
import NewestPosts from "@/components/home/NewestPosts.vue";
import Contact from "@/components/home/Contact.vue"
</script>

<template>
  <div> 
    <Topics />
    <div class="divider"></div> 
    <NewestPosts />

    <div class="large-space"></div>

    <section id="home-discussions">
      <div class="wrapper">
        <h4 class="font-size-title wrap">
          <span class="orange-text">Trending</span>
          Discussions
        </h4>
        <div class="grid">
          <article class="glow s12 m6 l6 no-margin" v-for="(i, x) in 6">
            <h6>Discussion #{{ i }}</h6>
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

    <div class="space"></div>
    <div class="divider"></div>
    <div class="space"></div>
    <Contact /> 
  </div>
</template>
