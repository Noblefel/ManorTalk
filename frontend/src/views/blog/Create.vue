<script setup lang="ts">
import Markdown from "@/components/Markdown.vue";
import { onMounted, ref } from "vue";
import {
  type CreatePost,
  usePostStore,
  type Category,
  sampleContent,
} from "@/stores/post";
import { RequestResponse } from "@/utils/api";
import { onBeforeRouteUpdate } from "vue-router";

const postStore = usePostStore();
const rr = ref(new RequestResponse());
const rr2 = ref(new RequestResponse());

const form = ref<CreatePost>({
  title: "",
  excerpt: "",
  content: sampleContent,
  category_id: 1,
});

const render = ref(false);

onMounted(() => {
  postStore.fetchCategories(rr2.value);
});
</script>

<template>
  <form class="grid" @submit.prevent="postStore.createPost(form, rr)">
    <div class="s12 m12 l9">
      <h3 class="center-align">Create Post 🎨</h3>
      <div class="space"></div>
      <label for="title" class="font-size-1-25 font-600">Title</label>
      <div class="field border no-margin">
        <input type="text" id="title" name="title" v-model="form.title" />
        <span class="error" v-if="rr.errors?.title">
          {{ rr.errors.title[0] }}
        </span>
      </div>

      <div class="space"></div>

      <label for="excerpt" class="font-size-1-25 font-600">Excerpt</label>
      <div class="field border no-margin">
        <input type="text" id="excerpt" name="excerpt" v-model="form.excerpt" />
        <span class="error" v-if="rr.errors?.excerpt">
          {{ rr.errors.excerpt[0] }}
        </span>
      </div>

      <div class="space"></div>

      <label for="category" class="font-size-1-25 font-600">Category</label>
      <div v-if="rr2.data" class="field border no-margin suffix">
        <select name="category" id="category" v-model="form.category_id">
          <option v-for="c in (rr2.data as any as Category[])" :value="c.id">
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
      <div class="actions">
        <button
          type="button"
          class="inverted responsive"
          @click="
            () => {
              form.title = '';
              form.excerpt = '';
              form.content = sampleContent;
              form.category_id = 1;
            }
          "
        >
          Reset
          <i>cancel</i>
        </button>
        <div class="space"></div>
        <button
          type="submit"
          class="secondary responsive"
          :disabled="rr.loading"
        >
          {{ rr.loading ? "Processing..." : "Done" }}
          <i v-if="!rr.loading">done</i>
          <progress v-else class="circle small"></progress>
        </button>
      </div>
    </div>
  </form>
</template>

<style scoped>
form {
  margin: auto;
  max-width: 1000px;
  padding: 1rem;
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
