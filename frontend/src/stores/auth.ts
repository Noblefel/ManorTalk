import { Api, RequestResponse } from "@/utils/api";
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
  const remember = ref(false);
  
  const isAuth = computed(() => authUser.value != null);

  function reset() {
    authUser.value = null;
    remember.value = false;
    localStorage.removeItem("user");
    localStorage.removeItem("access_token");
    sessionStorage.removeItem("user");
    sessionStorage.removeItem("access_token");
  }

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

  function setAuthStorage(access_token: string = "") {
    const userString = JSON.stringify(authUser.value);

    remember.value
      ? localStorage.setItem("user", userString)
      : sessionStorage.setItem("user", userString);

    if (access_token == "") return;

    remember.value
      ? localStorage.setItem("access_token", access_token)
      : sessionStorage.setItem("access_token", access_token);
  }

  async function login(form: LoginForm, rr: RequestResponse) {
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

    rr.loading = true;
    try {
      const res = await Api.post("/auth/login", form);
      const { access_token, user } = res.data.data;

      authUser.value = user;
      remember.value = form.remember_me;
      setAuthStorage(access_token);
      router.push({ name: "home" });
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  async function register(form: RegisterForm, rr: RequestResponse) {
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

    rr.loading = true;
    try {
      const res = await Api.post("/auth/register", form);
      toast(res.data.message, "green white-text");
      router.push({ name: "login" });
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  async function logout() {
    Api.post("/auth/logout", null);
    reset();
    window.location.reload();
  }

  return {
    authUser,
    remember,
    isAuth,
    reset,
    getAuthStorage,
    setAuthStorage,
    login,
    register,
    logout,
  };
});
