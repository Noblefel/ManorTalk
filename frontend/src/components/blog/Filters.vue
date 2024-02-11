<script setup lang="ts">
import { usePostStore } from "@/stores/post";
import { RequestResponse } from "@/utils/api";
import { changeParam } from "@/utils/helper";
import { ref } from "vue";
import { useRouter } from "vue-router";

const rr = ref(new RequestResponse());
const postStore = usePostStore();
const router = useRouter();
const search = ref("");

defineProps({ loading: Boolean });
</script>

<template>
  <div class="padding">
    <div class="field border prefix">
      <i>search</i>
      <input
        type="text"
        name="search"
        id="search"
        placeholder="Search"
        @input="(e) => { search = (e.target as HTMLInputElement).value }"
        :value="search"
        @keydown.enter="changeParam(router, 'search', search)"
        :disabled="loading"
      />
    </div>
    <div class="row wrap">
      <button class="inverted" :disabled="loading">
        <i>calendar_month</i>
        {{ $route.query.order == "asc" ? "Oldest" : "Newest" }}
        <i>arrow_drop_down</i>
        <menu>
          <a @click="changeParam(router, 'order', 'desc')">Newest</a>
          <a @click="changeParam(router, 'order', 'asc')">Oldest</a>
        </menu>
      </button>

      <button
        class="inverted"
        @click="postStore.fetchCategories(rr)"
        :disabled="loading"
      >
        <i>category</i>
        <span class="capitalize">
          {{
            (!$route.query?.category ? null : $route.query?.category) ??
            "All Categories"
          }}
        </span>
        <i>arrow_drop_down</i>

        <menu class="center-align" v-if="rr.errors || rr.loading">
          <div class="space"></div>
          <i>error</i>
          <div class="space"></div>
        </menu>
        <menu v-else>
          <a @click="changeParam(router, 'category', '')">All</a>
          <a
            v-for="c in postStore.categories"
            :key="c.slug"
            @click="changeParam(router, 'category', c.slug)"
          >
            {{ c.name }}
          </a>
        </menu>
      </button>

      <button class="inverted" :disabled="loading">
        <i>list_alt</i>
        {{ $route.query.limit ?? 10 }} per Page
        <i>arrow_drop_down</i>
        <menu>
          <a @click="changeParam(router, 'limit', 10)">10 per Page</a>
          <a @click="changeParam(router, 'limit', 20)">20 per Page</a>
          <a @click="changeParam(router, 'limit', 50)">50 per Page</a>
        </menu>
      </button>
    </div>
  </div>
</template>

<style scoped>
button {
  padding: 0 0.8rem;
  font-size: 0.75rem;

  & i {
    font-size: 1.2rem;
  }
}

menu {
  background-color: var(--background);
  border: 2px solid var(--primary);
  min-width: 10rem;

  a {
    transition: 0.1s all ease-in;

    &:hover {
      background-color: unset;
      color: var(--primary);
    }
  }
}
</style>
