<script setup lang="ts">
import type { Post } from "@/stores/post";
import { getAvatar } from "@/utils/helper";
import type { PropType } from "vue";
import { RouterLink } from "vue-router";

defineProps({
  post: {
    type: Object as PropType<Post>,
    default: {
      id: 0,
      title: "A sample post 0",
      slug: "a-sample-post-0",
      category: { name: "Sample" },
      user: { username: "john-doe" },
    },
  },
  width: {
    type: String,
    default: "100%",
  },
  image: {
    type: String,
  },
  separateImage: {
    type: Boolean,
    default: false,
  },
  imageHeight: {
    type: String,
    default: "clamp(15rem, calc(8rem + 8vw), 20rem)",
  },
  backgroundColor: {
    type: String,
    default: "var(--surface)",
  },
  withBorder: {
    type: Boolean,
    default: false,
  },
  withExcerpt: {
    type: Boolean,
    default: false,
  },
  withAuthor: {
    type: Boolean,
    default: false,
  },
  withStats: {
    type: Boolean,
    default: false,
  },
});

function cropText(text: String) {
  let max = 100;

  if (window.screen.width <= 600) {
    max = 50;
  } else if (window.screen.width <= 900) {
    max = 75;
  }

  return text.slice(0, max) + "...";
}
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
    <div v-if="image && !separateImage" class="merged-image">
      <img
        :src="image"
        alt=""
        :style="{
          height: imageHeight,
        }"
      />
      <div class="title-absolute">
        <p class="font-size-1-25">
          {{ cropText(post.title) }}
        </p>
        <div class="tags font-400">
          <div>Business</div>
        </div>
      </div>
    </div>
    <div v-else class="separate-image">
      <div v-if="image">
        <img
          :src="image"
          :alt="post.title"
          :style="{
            height: imageHeight,
          }"
        />
      </div>
      <div class="space"></div>
      <div class="tags font-400">
        <div>{{ post.category.name ?? "tes" }}</div>
      </div>
      <p class="font-size-1-25">
        {{ post.title }}
      </p>
      <div class="divider"></div>
    </div>
    <div class="details">
      <p class="excerpt" v-if="withExcerpt">
        <!-- {{ post.excerpt }} -->
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
        sapien luctus libero ultricies egestas vel vitae ligula. Aliquam posuere
        condimentum ex, quis luctus libero interdum eget"
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
          <p class="small-text font-600" v-if="post.created_at">
            {{ new Date(post.created_at).toUTCString() }}
          </p>
        </div>
      </div>
      <div class="stats" v-if="withStats">
        <div class="green1 green-text small-opacity">
          <i>thumb_up</i>
          <span>12</span>
        </div>
        <div class="secondary">
          <i>comment</i>
          <span>0</span>
        </div>
        <div class="red2 red-text small-opacity">
          <i>warning</i>
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
  padding: 0.5rem 1rem;
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

  p {
    margin: 0;
  }

  .stats {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    font-weight: 600;

    & > * {
      display: flex;
      gap: 0.5rem;
      margin: 0;
      padding: 0.2rem 0.5rem;
      border-radius: 7px;
      height: fit-content;
      font-size: 0.9rem;

      & i {
        font-size: 1.2rem;
      }
    }
  }
}
</style>
