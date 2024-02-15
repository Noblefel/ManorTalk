<script setup lang="ts">
import AuthCard from "@/components/auth/AuthCard.vue";
import { ref } from "vue";
import { RequestResponse } from "@/utils/api";
import { useAuthStore } from "@/stores/auth";
import { useUserStore, type UpdateForm } from "@/stores/user";
import { getAvatar } from "@/utils/helper";

const authStore = useAuthStore();
const userStore = useUserStore();

const u = authStore.authUser;

const form = ref<UpdateForm>({
  name: u?.name,
  username: u?.username ?? "",
  avatar: u?.avatar,
  bio: u?.bio,
});

const rr = ref(new RequestResponse());
const rrCheck = ref(new RequestResponse());
</script>

<template>
  <AuthCard title="Edit Profile 📝">
    <form @submit.prevent="userStore.update(form, rr, $route.params.username)">
      <div class="space"></div>

      <div class="center-align">
        <img :src="getAvatar(u)" alt="" width="75px" />
      </div>

      <div class="padding">
        <label for="email" class="font-size-0-9 font-600">Name</label>
        <div class="field border no-margin prefix">
          <i>account_circle</i>
          <input
            type="text"
            name="name"
            id="name"
            autocomplete="off"
            v-model.trim="form.name"
          />
          <span class="error" v-if="rr.errors?.name">
            {{ rr.errors.name[0] }}
          </span>
        </div>

        <div class="space"></div>

        <label for="email" class="font-size-0-9 font-600">Username</label>
        <div class="field border no-margin prefix">
          <i>tag</i>
          <input
            type="text"
            name="username"
            id="username"
            autocomplete="off"
            v-model.trim="form.username"
          />
          <i
            class="cursor-pointer z-2"
            @click="
              () => {
                if (u?.username == form.username) return;
                userStore.checkUsername(form.username, rrCheck);
              }
            "
            v-if="!rrCheck.loading"
          >
            search
          </i>
          <i v-else>
            <progress class="circle surface small"></progress>
          </i>
          <span
            class="error"
            v-if="rr.errors?.username || rrCheck.errors?.username"
          >
            {{ rr.errors?.username[0] || rrCheck.errors?.username[0] }}
          </span>
        </div>

        <div class="space"></div>

        <label for="avatar" class="font-size-0-9 font-600">Avatar</label>
        <div class="field border no-margin prefix">
          <i>attach_file</i>
          <input type="file" />
          <input
            type="text"
            id="avatar"
            name="avatar"
            placeholder="Click to change"
          />
          <span class="error" v-if="rr.errors?.bio">
            {{ rr.errors.bio[0] }}
          </span>
        </div>

        <div class="space"></div>

        <label for="bio" class="font-size-0-9 font-600">Bio</label>
        <div class="field border no-margin prefix">
          <i>article</i>
          <input
            type="text"
            name="bio"
            id="bio"
            autocomplete="off"
            v-model.trim="form.bio"
          />
          <span class="error" v-if="rr.errors?.bio">
            {{ rr.errors.bio[0] }}
          </span>
        </div>
      </div>

      <div class="space"></div>

      <div class="row center-align wrap">
        <RouterLink
          :to="{ name: 'profile', params: $route.params }"
          class="button secondary"
        >
          Cancel
          <i>cancel</i>
        </RouterLink>
        <button :disabled="rr.loading">
          {{ rr.loading ? "Updating..." : "Update" }}
          <i v-if="!rr.loading">edit</i>
          <progress v-else class="circle white-text small"></progress>
        </button>
      </div>
    </form>
  </AuthCard>
</template>

<style scoped>
p {
  font-size: 0.9rem;
}

img {
  border-radius: 50%;
}
</style>
