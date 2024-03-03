<script setup lang="ts">
import Markdown from "@/components/Markdown.vue";
import { onMounted, ref } from "vue";
import { type CreatePost, usePostStore, sampleContent } from "@/stores/post";
import { RequestResponse } from "@/utils/api";
import { verifyImage } from "@/utils/helper";
import Actions from "@/components/blog/Actions.vue";

const postStore = usePostStore();
const rr = ref(new RequestResponse());
const rr2 = ref(new RequestResponse());

const form = ref<CreatePost>({
  title: "",
  excerpt: "",
  content: sampleContent,
  category_id: 1,
  image: null,
});

const render = ref(false);

onMounted(() => {
  postStore.fetchCategories(rr2.value);
});

const shownImage = ref("");

function onFileChange(event: Event) {
  form.value.image = null;
  const files = (event.target as HTMLInputElement).files;
  const img = verifyImage(files);
  if (img && files) {
    form.value.image = files[0];
    shownImage.value = img;
  }
}
</script>

<template>
  <form class="grid" @submit.prevent="postStore.create(form, rr)">
    <div class="s12 m12 l9">
      <h3 class="center-align">Create Post 🎨</h3>

      <div class="space"></div>
      <img
        v-if="form.image"
        :src="shownImage"
        alt="Post Image"
        class="responsive medium-height small-round"
      />
      <div class="space"></div>

      <label for="title" class="font-size-1-25 font-600">Title</label>
      <div class="field border no-margin">
        <input
          type="text"
          id="title"
          name="title"
          v-model="form.title"
          autocomplete="off"
        />
        <span class="error" v-if="rr.errors?.title">
          {{ rr.errors.title[0] }}
        </span>
      </div>

      <div class="space"></div>

      <label for="excerpt" class="font-size-1-25 font-600">Excerpt</label>
      <div class="field border no-margin">
        <input
          type="text"
          id="excerpt"
          name="excerpt"
          v-model="form.excerpt"
          autocomplete="off"
        />
        <span class="error" v-if="rr.errors?.excerpt">
          {{ rr.errors.excerpt[0] }}
        </span>
      </div>

      <div class="space"></div>

      <label for="category" class="font-size-1-25 font-600">Category</label>
      <div v-if="postStore.categories" class="field border no-margin suffix">
        <select name="category" id="category" v-model="form.category_id">
          <option v-for="c in postStore.categories" :value="c.id">
            {{ c.name }}
          </option>
        </select>
        <i>arrow_drop_down</i>
      </div>

      <div class="space"></div>

      <label for="content" class="font-size-1-25 font-600">Content</label>
      <div class="row no-space no-margin">
        <button
          type="button"
          class="left-round border"
          :class="{ secondary: !render }"
          @click="render = false"
        >
          <i>edit</i>Markdown
        </button>
        <button
          type="button"
          class="right-round border"
          :class="{ secondary: render }"
          @click="render = true"
        >
          <i>article</i>Render
        </button>
      </div>
      <span class="red-text" v-if="rr.errors?.content">
        <div class="space"></div>
        {{ rr.errors.content[0] }}
      </span>
      <Markdown v-if="render" :content="form.content" />
      <div v-else class="field border textarea large-height">
        <textarea
          type="text"
          id="content"
          name="content"
          v-model="form.content"
        >
        </textarea>
      </div>
    </div>
    <div class="s12 m12 l3">
      <Actions
        :on-file-change="onFileChange"
        :on-reset="
          () => {
            form.title = '';
            form.excerpt = '';
            form.content = sampleContent;
            form.category_id = 1;
            form.image = null;
          }
        "
        :rr-submit="rr"
      />
    </div>
  </form>
</template>

<style scoped>
form {
  margin: auto;
  max-width: 1000px;
  padding: 1rem;
}

img {
  object-fit: cover;
}

textarea,
textarea:focus {
  border: 2px dashed;
}

@media screen and (min-width: 992px) {
  .actions {
    position: fixed;
    width: 15rem;
  }
}
</style>
