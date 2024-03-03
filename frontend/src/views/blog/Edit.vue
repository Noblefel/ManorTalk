<script setup lang="ts">
import Markdown from "@/components/Markdown.vue";
import ResponseCard from "@/components/ResponseCard.vue";
import Actions from "@/components/blog/Actions.vue";
import { onMounted, ref } from "vue";
import { type CreatePost, usePostStore } from "@/stores/post";
import { RequestResponse } from "@/utils/api";
import { useRoute } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { getImage, verifyImage } from "@/utils/helper";

const postStore = usePostStore();
const authStore = useAuthStore();
const rr = ref(new RequestResponse()); // form
const rr2 = ref(new RequestResponse()); // categories
const rr3 = ref(new RequestResponse()); // viewedPost
const rr4 = ref(new RequestResponse()); // delete
const route = useRoute();
const form = ref({} as CreatePost);
const shownImage = ref("");

function reset() {
  form.value.title = postStore.viewedPost?.title ?? "";
  form.value.excerpt = postStore.viewedPost?.excerpt ?? "";
  form.value.content = postStore.viewedPost?.content ?? "";
  form.value.category_id = postStore.viewedPost?.category_id ?? 1;
  form.value.image = null;
  shownImage.value = getImage("post/" + postStore.viewedPost?.image ?? "");
}

function onFileChange(event: Event) {
  form.value.image = null;
  const files = (event.target as HTMLInputElement).files;
  const img = verifyImage(files);
  if (img && files) {
    form.value.image = files[0];
    shownImage.value = URL.createObjectURL(files[0]);
  }
}

const render = ref(false);

onMounted(() => {
  postStore.fetchPost(rr3.value, route).then(() => {
    reset();
  });
  postStore.fetchCategories(rr2.value);
});
</script>

<template>
  <form @submit.prevent="postStore.update(form, rr)">
    <ResponseCard
      v-if="
        postStore.viewedPost &&
        postStore.viewedPost?.user_id != authStore.authUser?.id
      "
      message="You have no permission to view this"
      icon="warning"
    />
    <div class="grid" v-else-if="postStore.viewedPost">
      <div class="s12 m12 l9">
        <h3 class="center-align">Edit Post 🎨</h3>

        <div class="space"></div>
        <img
          v-if="shownImage"
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
          :on-reset="reset"
          :on-delete="postStore.deletePost"
          :rr-delete="rr4"
          :rr-submit="rr"
        />
      </div>
    </div>
    <ResponseCard
      v-else-if="rr3.errors"
      :message="`Unable to get post (${rr3.status})`"
      icon="error"
    />
    <ResponseCard v-else :loading="true" />
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
</style>
