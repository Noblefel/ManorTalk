import type { RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { type RouteLocation, useRouter } from "vue-router";
import { useAuthStore } from "./auth";
import { ref } from "vue";

export interface User {
  id: number;
  name?: string;
  username: string;
  avatar: string;
  email: string;
  password: string;
  created_at?: string;
  updated_at?: string;
  posts_count?: number;
  bio?: string;
}

export interface UpdateForm {
  name?: string;
  username: string;
  avatar: File | null;
  bio?: string;
}

export const useUserStore = defineStore("user", () => {
  const router = useRouter();
  const authStore = useAuthStore();
  const viewedUser = ref<User | null>(null);

  /** checkUsername validates the username and send request to check its availability */
  function checkUsername(username: string, rr: RequestResponse) {
    const f = new Validator({ username: username })
      .required("username")
      .strMinLength("username", 3)
      .strMaxLength("username", 40);

    if (!f.isValid()) {
      rr.errors = f.errors;
      return;
    }

    rr.useApi("post", "/users/check-username", f.form).then(() => {
      if (rr.status != 200) return;
      rr.errors = null;

      if (rr.message) toast(rr.message, "green white-text");
    });
  }

  /** fetchProfile will get the user profile data */
  function fetchProfile(to: RouteLocation, rr: RequestResponse) {
    if (viewedUser.value?.username == to.params.username) {
      return Promise.resolve();
    }

    if (to.params.username == authStore.authUser?.username) {
      viewedUser.value = authStore.authUser;
      return Promise.resolve();
    }

    return rr.useApi("GET", "/users/" + to.params.username).then(() => {
      if (rr.status !== 200) return;

      viewedUser.value = rr.data as any;
    });
  }

  /** update validates the form and updates the user profile */
  function update(form: UpdateForm, rr: RequestResponse, username: string) {
    const f = new Validator(form)
      .required("username")
      .strMinLength("username", 3)
      .strMaxLength("username", 40)
      .strMaxLength("name", 255)
      .strMaxLength("bio", 2000);

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    }

    rr.useApi(
      "patch",
      `/users/${username}`,
      form,
      true,
      "multipart/form-data"
    ).then(() => {
      if (rr.status != 200) return;

      authStore!.authUser!.name = form.name;
      authStore!.authUser!.username = form.username;
      authStore!.authUser!.bio = form.bio;
      authStore!.authUser!.avatar = rr.data as any;
      authStore.setAuthStorage();

      if (rr.message) toast(rr.message, "green white-text");

      router.push({ name: "profile", params: { username: form.username } });
    });
  }

  return {
    viewedUser,
    checkUsername,
    fetchProfile,
    update,
  };
});
