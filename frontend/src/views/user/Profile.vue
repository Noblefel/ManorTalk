<script setup lang="ts">
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import Header from "@/components/user/Header.vue";
import type { User } from "@/stores/user";

const route = useRoute();
const rr = ref(new RequestResponse());
const fetchPosts = ref(false);

onMounted(() => {
  window.scrollTo(0, 0);
  rr.value.useApi("GET", "/user/" + route.params.username).then(() => {
    if (rr.value.status == 200) fetchPosts.value = true;
  });
});
</script>

<template>
  <Header
    :user="(rr.data as any as User)"
    :errors="rr.errors != null"
    :loading="rr.loading"
    :status="rr.status"
  >
    <div class="grid">
      <div class="s12 m12 l4">
        <article>
          <div class="row">
            <div class="chip circle primary no-border large">
              <p class="font-size-1-5">📋</p>
            </div>
            <p class="font-600">About Me</p>
          </div>

          <div class="item">
            <i>person</i>
            <p>Full Name:</p>
            <p>{{ (rr.data as any).name }}</p>
          </div>
          <div class="item">
            <i>account_circle</i>
            <p>Username:</p>
            <p>{{ (rr.data as any).username }}</p>
          </div>
          <div class="item">
            <i>article</i>
            <p>Post Created:</p>
            <p>{{ (rr.data as any).posts_count ?? 0 }} posts</p>
          </div>
          <div class="item">
            <i>calendar_month</i>
            <p>Joined:</p>
            <p>{{ new Date((rr.data as any).created_at).toUTCString() }}</p>
          </div>
        </article>
      </div>
      <div class="s12 m12 l8">
        <article>
          <div class="row">
            <div class="chip circle primary no-border large">
              <p class="font-size-1-5">📜</p>
            </div>
            <p class="font-600">Bio</p>
          </div>

          <div
            v-if="!(rr.data as any).bio"
            class="center-align small-opacity font-500"
          >
            <div class="large-space"></div>
            <i>face</i>
            <p>No Bio</p>
            <div class="large-space"></div>
          </div>
        </article>
      </div>
    </div>
  </Header>
</template>

<style scoped>
article {
  margin: 0;
  background-color: var(--background);
}

.chip {
  margin: 0;
}

.item {
  display: flex;
  gap: 0.5rem;
  font-size: 1rem;
  flex-wrap: wrap;
  margin-top: 1rem;

  p:first-of-type {
    font-weight: 600;
    margin: 0;
  }

  p:last-of-type {
    opacity: 0.9;
    margin: 0rem;
  }
}
</style>
