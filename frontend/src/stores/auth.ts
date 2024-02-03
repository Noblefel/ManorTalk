// import { useApi } from "@/utils/api";
import type { RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";

export interface User {
  email: string;
  password: string;
  created_at?: string;
  updated_at?: string;
}

export interface LoginForm {
  email: string;
  password: string;
  remember_me: boolean;
}

export interface RegisterForm {
  email: string;
  password: string;
  password_repeat: string;
}

export const useAuthStore = defineStore("auth", () => {
  const router = useRouter();

  const authUser = ref<User | null>();

  /** isAuth checks if authUser store is not null */
  const isAuth = computed(() => authUser != null);

  /** login validates the form and attempts to authenticate the user */
  function login(form: LoginForm, rr: RequestResponse) {
    const f = new Validator(form)
      .required("email", "password")
      .email("email")
      .strMinLength("password", 8)
      .strMaxLength("password", 255);

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
      form.remember_me
        ? localStorage.setItem("access_token", access_token)
        : sessionStorage.setItem("access_token", access_token);

      router.push({ name: "home" });
    });
  }

  /** register validates the form and creates new user */
  function register(form: RegisterForm, rr: RequestResponse) {
    const f = new Validator(form)
      .required("email", "password")
      .email("email")
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

  return {
    authUser,
    isAuth,
    login,
    register,
  };
});
