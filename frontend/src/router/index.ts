import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import AppLayout from "@/components/layout/AppLayout.vue"
import AuthLayout from "@/components/layout/AuthLayout.vue"

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
      meta: { layout: AuthLayout },
      component: () => import("@/views/auth/Login.vue"),
    },
    {
      path: "/register",
      name: "register",
      meta: { layout: AuthLayout },
      component: () => import("@/views/auth/Register.vue"),
    },
  ],
});

export default router;
