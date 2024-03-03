<script setup lang="ts">
import type { Post } from "@/stores/post";
import { getAvatar, getImage } from "@/utils/helper";
import type { PropType } from "vue";
import { RouterLink } from "vue-router";

defineProps({
  post: { type: Object as PropType<Post>, required: true },
  width: { type: String, default: "100%" }, 
  imageHeight: {
    type: String,
    default: "clamp(15rem, calc(8rem + 8vw), 20rem)",
  },
  backgroundColor: { type: String, default: "var(--surface)" },
  withBorder: { type: Boolean, default: false },
  withExcerpt: { type: Boolean, default: false },
  withAuthor: { type: Boolean, default: false },
});
</script>

<template>
  <article
    :class="{
      'no-border': !withBorder,
      glow: withBorder,
    }"
    :style="{
      'background-color': backgroundColor,
      width: width,
    }"
  >
    <div class="separate-image">
      <div v-if="post.image">
        <img
          :src="getImage(`post/${post.image}`)"
          :alt="post.title"
          :style="{
            height: imageHeight,
          }"
          loading="lazy"
        />
      </div>
      <div class="space"></div>
      <div class="tags font-400">
        <div>{{ post.category.name ?? "tes" }}</div>
      </div>
      <RouterLink
        :to="{ name: 'blog.show', params: { slug: post.slug } }"
        class="font-size-1-25"
      >
        {{ post.title }}
      </RouterLink>
    </div>
    <div class="details">
      <p class="excerpt" v-if="withExcerpt">
        {{ post.excerpt }}
      </p>
      <div class="author" v-if="withAuthor">
        <img
          :src="getAvatar(post.user)"
          :alt="`${post.user.username} profile avatar`"
        />
        <div>
          <RouterLink
            :to="{ name: 'profile', params: { username: post.user.username } }"
            class="font-700"
          >
            {{ post.user.name ?? post.user.username }}
          </RouterLink>
          <p class="small-text font-600 no-margin" v-if="post.created_at">
            {{ new Date(post.created_at).toUTCString() }}
          </p>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
article {
  padding: 0 0 0.5rem 0;
  margin: 0;
}

.merged-image {
  img {
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
    width: 100%;
    object-fit: cover;
    filter: brightness(0.75);
  }

  .title-absolute {
    padding: 0.5rem 1rem;
    font-weight: 600;
    position: absolute;
    bottom: 0;
    color: white;
  }
}

.separate-image {
  padding: 1rem;
  font-weight: 600;
  margin-bottom: -0.5rem;

  img {
    width: 100%;
    object-fit: cover;
  }

  p {
    opacity: 0.9;
  }
}

.tags {
  display: flex;
  flex-wrap: wrap;
  font-size: 0.9rem;
  gap: 0.5rem;
  color: white;

  > * {
    background-image: linear-gradient(to right, rgb(56, 53, 223), #8f24d6);
    padding: 0.1rem 0.5rem;
    border-radius: 5px;
  }
}

.details {
  padding: 0 1rem 0.5rem 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;

  & .excerpt {
    font-size: 0.95rem;
    opacity: 0.9;
  }

  & .author {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;

    img {
      width: 40px;
      height: 40px;
      object-fit: cover;
      border-radius: 50%;
    }
  }
}
</style>
