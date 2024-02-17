<script setup lang="ts">
import AuthCard from "@/components/auth/AuthCard.vue";
import { ref } from "vue";
import { RequestResponse } from "@/utils/api";
import { useAuthStore } from "@/stores/auth";
import { useUserStore, type UpdateForm } from "@/stores/user";
import { getAvatar, toast } from "@/utils/helper";

const authStore = useAuthStore();
const userStore = useUserStore();

const u = authStore.authUser;

const form = ref<UpdateForm>({
  name: u?.name,
  username: u?.username ?? "",
  avatar: null,
  bio: u?.bio,
});

const rr = ref(new RequestResponse());
const rrCheck = ref(new RequestResponse());

const shownImage = ref(getAvatar(u));

function onFileChange(event: Event) {
  form.value.avatar = null;
  shownImage.value = getAvatar(u);
  const file = (event.target as HTMLInputElement).files;

  if (!file) {
    toast("File must be an image");
    return;
  }

  if (file[0].size > 2 * 1024 * 1024) {
    toast("Image is too large (2MB max)");
    return;
  }

  form.value.avatar = file[0];
  shownImage.value = URL.createObjectURL(file[0]);
}
</script>

<template>
  <AuthCard title="Edit Profile 📝">
    <form
      @submit.prevent="
        userStore.update(form, rr, $route.params.username as string)
      "
      enctype="multipart/form-data"
    >
      <div class="space"></div>

      <div class="center-align">
        <img :src="shownImage" alt="Avatar" width="75px" height="75px" />
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

        <label for="username" class="font-size-0-9 font-600">Username</label>
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
          <input type="file" @change="onFileChange" accept=".jpeg,.jpg,.png" />
          <input
            type="text"
            id="avatar"
            name="avatar"
            placeholder="Click to change"
          />
        </div>

        <div class="space"></div>
        <label class="font-size-0-9 font-600">Bio</label>
        <button
          type="button"
          class="responsive secondary no-margin"
          data-ui="#edit-bio"
        >
          <i>article</i>
          View Bio
        </button>
        <span class="error-text font-size-0-9" v-if="rr.errors?.bio">
          {{ rr.errors.bio[0] }}
        </span>

        <dialog id="edit-bio">
          <h6>My Bio 📓</h6>
          <div class="field textarea no-border">
            <textarea id="bio" name="bio" v-model="form.bio"></textarea>
          </div>
          <div class="row right-align">
            <button class="secondary" type="button" @click="form.bio = u?.bio">
              <i>undo</i>
              Undo
            </button>
            <button class="secondary" data-ui="#edit-bio" type="button">
              Ok
            </button>
          </div>
        </dialog>
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
        <button type="submit" :disabled="rr.loading">
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

#edit-bio {
  width: 700px;
  border-radius: 8px;
  background-color: var(--background);
  border: 1px solid var(--secondary);
  transition: none;

  .field {
    block-size: 25rem;
  }

  textarea {
    border: 2px dashed;
  }
}
</style>
