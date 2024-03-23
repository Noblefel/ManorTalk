import { Api, type RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { type RouteLocation, useRouter } from "vue-router";
import { useAuthStore } from "./auth";
import { ref, type Ref } from "vue";
import type { Post } from "./post";

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
  posts?: Post[];
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

  async function checkUsername(username: string, rr: RequestResponse) {
    const f = new Validator({ username: username })
      .required("username")
      .strMinLength("username", 3)
      .strMaxLength("username", 40);

    if (!f.isValid()) {
      rr.errors = f.errors;
      return;
    }

    rr.loading = true;
    try {
      const res = await Api.post("/users/check-username", f.form);
      rr.errors = null;
      toast(res.data.message, "green white-text");
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  async function fetchProfile(rr: Ref<RequestResponse>, to: RouteLocation) {
    if (viewedUser.value?.username == to.params.username) {
      return;
    }

    if (authStore.authUser?.username == to.params.username) {
      viewedUser.value = authStore.authUser;
      return;
    }

    rr.value.loading = true;
    try {
      const res = await Api.get("/users/" + to.params.username);
      viewedUser.value = res.data.data;
    } catch (e) {
      rr.value.handleErr(e);
    } finally {
      rr.value.loading = false;
    }
  }

  async function update(
    form: UpdateForm,
    rr: RequestResponse,
    to: RouteLocation
  ) {
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

    rr.loading = true;
    try {
      const res = await Api.patch("/users/" + to.params.username, form, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      authStore!.authUser!.name = form.name;
      authStore!.authUser!.username = form.username;
      authStore!.authUser!.bio = form.bio;
      authStore!.authUser!.avatar = res.data.data;
      authStore.setAuthStorage();

      toast(res.data.message, "green white-text");
      router.push({ name: "profile", params: { username: form.username } });
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  return {
    viewedUser,
    checkUsername,
    fetchProfile,
    update,
  };
});
