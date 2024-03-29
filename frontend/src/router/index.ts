import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import AppLayout from "@/components/layout/AppLayout.vue";
import AuthLayout from "@/components/layout/AuthLayout.vue";
import { useAuthStore } from "@/stores/auth";
import { guest, auth, checkAuth, userGuard } from "./middleware";

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
    {
      path: "/blog/create",
      name: "blog.create",
      meta: { layout: AppLayout, auth: true },
      component: () => import("@/views/blog/Create.vue"),
    },
    {
      path: "/blog/:slug",
      name: "blog.show",
      meta: { layout: AppLayout },
      component: () => import("@/views/blog/Show.vue"),
    },
    {
      path: "/blog/:slug/edit",
      name: "blog.edit",
      meta: { layout: AppLayout, auth: true },
      component: () => import("@/views/blog/Edit.vue"),
    },
    {
      path: "/profile/:username",
      name: "profile",
      meta: { layout: AppLayout },
      component: () => import("@/views/user/Profile.vue"),
    },
    {
      path: "/profile/:username/posts",
      name: "profile.posts",
      meta: { layout: AppLayout },
      component: () => import("@/views/user/Posts.vue"),
    },
    {
      path: "/profile/:username/edit",
      name: "profile.edit",
      meta: { layout: AuthLayout, auth: true, userGuard: true },
      component: () => import("@/views/user/Edit.vue"),
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();
  const ctx = { to, from, next, authStore };

  checkAuth(ctx);

  if (to.meta.guest && !guest(ctx)) return;

  if (to.meta.auth && !auth(ctx)) return;

  if (to.meta.userGuard && !userGuard(ctx)) return;

  if (to.name == "logout") {
    authStore.logout();
    return;
  }

  next();
});

export default router;
