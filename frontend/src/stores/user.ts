import type { RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { type RouteLocation, useRouter } from "vue-router";
import { useAuthStore } from "./auth";

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
  function fetchProfile(
    to: RouteLocation,
    rr: RequestResponse,
    authUser: User | null
  ) {
    window.scrollTo(0, 0);

    if (to.params.username == authUser?.username) {
      rr.data = authUser as any;
      return;
    }

    rr.useApi("GET", "/users/" + to.params.username);
  }

  /** update validates the form and updates the user profile */
  function update(form: UpdateForm, rr: RequestResponse, username: string) {
    const f = new Validator(form)
      .required("name", "username")
      .strMinLength("username", 3)
      .strMaxLength("username", 40)
      .strMaxLength("name", 255);

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
      authStore!.authUser!.avatar = rr.data as any;
      authStore.setAuthStorage();

      if (rr.message) toast(rr.message, "green white-text");

      router.push({ name: "profile", params: { username: form.username } });
    });
  }

  return {
    checkUsername,
    fetchProfile,
    update,
  };
});
