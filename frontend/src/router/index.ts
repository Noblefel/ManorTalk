import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import AppLayout from "@/components/layout/AppLayout.vue";
import AuthLayout from "@/components/layout/AuthLayout.vue";
import { useAuthStore } from "@/stores/auth";
import { guest, auth, checkAuth } from "./middleware";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      meta: { layout: AppLayout },
      component: Home,
    },
    {
      path: "/login",
      name: "login",
      meta: { layout: AuthLayout, guest: true },
      component: () => import("@/views/auth/Login.vue"),
    },
    {
      path: "/register",
      name: "register",
      meta: { layout: AuthLayout, guest: true },
      component: () => import("@/views/auth/Register.vue"),
    },
    {
      path: "/logout",
      name: "logout",
      component: Home, // Placeholder
    },
    {
      path: "/blog",
      name: "blog",
      meta: { layout: AppLayout },
      component: () => import("@/views/blog/Index.vue"),
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();
  const ctx = { to, from, next, authStore };

  checkAuth(ctx);

  if (to.meta.guest && !guest(ctx)) return;

  if (to.meta.auth && !auth(ctx)) return;

  if (to.name == "logout") {
    authStore.logout();
    return;
  }

  next();
});

export default router;
