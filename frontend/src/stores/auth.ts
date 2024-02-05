import { RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import type { User } from "./user";

export interface LoginForm {
  email: string;
  password: string;
  remember_me: boolean;
}

export interface RegisterForm {
  username: string;
  email: string;
  password: string;
  password_repeat: string;
}

export const useAuthStore = defineStore("auth", () => {
  const router = useRouter();

  const authUser = ref<User | null>(null);

  /** isAuth checks if authUser store is not null */
  const isAuth = computed(() => authUser.value != null);

  /** getAuthStorage will retrieve authenticated user data from the local/session storage */
  function getAuthStorage() {
    let user = null;
    let userString = localStorage.getItem("user");
    let access_token = localStorage.getItem("access_token");
    let remember = true;

    if (!access_token) {
      userString = sessionStorage.getItem("user");
      access_token = sessionStorage.getItem("access_token");
      remember = false;
    }

    user = userString ? (JSON.parse(userString) as User) : null;
    return { user, access_token, remember };
  }

  /** setAuthStorage will save authenticated user data in the local/session storage */
  function setAuthStorage(access_token: string, remember: boolean) {
    const userString = JSON.stringify(authUser.value);

    remember
      ? localStorage.setItem("user", userString)
      : sessionStorage.setItem("user", userString);

    remember
      ? localStorage.setItem("access_token", access_token)
      : sessionStorage.setItem("access_token", access_token);
  }

  /** login validates the form and attempts to authenticate the user */
  function login(form: LoginForm, rr: RequestResponse) {
    const f = new Validator(form)
      .required("email", "password")
      .email("email")
      .strMinLength("password", 8)
      .strMaxLength("password", 72);

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    }

    rr.useApi("post", "/auth/login", form).then(() => {
      if (rr.status != 200) return;

      const { access_token, user } = rr.data as unknown as {
        access_token: string;
        user: User;
      };

      authUser.value = user;
      setAuthStorage(access_token, form.remember_me);
      router.push({ name: "home" });
    });
  }

  /** register validates the form and creates new user */
  function register(form: RegisterForm, rr: RequestResponse) {
    const f = new Validator(form)
      .required("email", "password", "username")
      .email("email")
      .strMinLength("username", 3)
      .strMaxLength("username", 40)
      .strMinLength("password", 8)
      .strMaxLength("password", 72)
      .equal("password_repeat", "password");

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    }

    rr.useApi("post", "/auth/register", form).then(() => {
      if (rr.status != 200) return;

      if (rr.message) toast(rr.message, "green white-text");

      router.push({ name: "login" });
    });
  }

  /** logout will reset the authentication state */
  function logout() {
    // removes the refresh token from cache
    const rr = new RequestResponse();

    rr.useApi("post", "/auth/logout", null, false);

    authUser.value = null;
    localStorage.removeItem("user");
    localStorage.removeItem("access_token");
    sessionStorage.removeItem("user");
    sessionStorage.removeItem("access_token");

    toast("Logged out", "green white-text");
  }

  return {
    authUser,
    isAuth,
    getAuthStorage,
    setAuthStorage,
    login,
    register,
    logout,
  };
});
