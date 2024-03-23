<script setup lang="ts">
import Header from "@/components/user/Header.vue";
import { useUserStore } from "@/stores/user";
import { RequestResponse } from "@/utils/api";
import { onMounted, ref } from "vue";
import { onBeforeRouteUpdate, useRoute } from "vue-router";
const as = useUserStore();

const rr = ref(new RequestResponse());
const route = useRoute();

onMounted(() => {
  as.fetchProfile(rr, route);
});

onBeforeRouteUpdate((to) => {
  as.fetchProfile(rr, to);
});
</script>

<template>
  <Header :rr="rr">
    <div class="grid">
      <div class="s12 m12 l4">
        <article v-if="as.viewedUser">
          <div class="row">
            <div class="chip circle primary no-border large">
              <p class="font-size-1-5">📋</p>
            </div>
            <p class="font-600">About Me</p>
          </div>

          <div class="item">
            <i>person</i>
            <p>Full Name:</p>
            <p>{{ as.viewedUser.name }}</p>
          </div>
          <div class="item">
            <i>account_circle</i>
            <p>Username:</p>
            <p>{{ as.viewedUser.username }}</p>
          </div>
          <div class="item">
            <i>article</i>
            <p>Post Created:</p>
            <p>{{ as.viewedUser.posts_count ?? 0 }} posts</p>
          </div>
          <div class="item">
            <i>calendar_month</i>
            <p>Joined:</p>
            <p>
              {{
                new Date(as.viewedUser.created_at || "").toUTCString()
              }}
            </p>
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
            v-if="!as.viewedUser?.bio"
            class="center-align small-opacity font-500"
          >
            <div class="large-space"></div>
            <i>face</i>
            <p>No Bio</p>
            <div class="large-space"></div>
          </div>
          <div v-else>
            <p v-for="p in as.viewedUser.bio.split('\n')">
              {{ p == "" ? "&nbsp;" : p }}
            </p>
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
