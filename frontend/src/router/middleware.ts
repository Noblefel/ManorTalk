import type { useAuthStore } from "@/stores/auth";
import type { NavigationGuardNext, RouteLocationNormalized } from "vue-router";

export interface MiddlewareCtx {
  to: RouteLocationNormalized;
  from: RouteLocationNormalized;
  next: NavigationGuardNext;
  authStore: ReturnType<typeof useAuthStore>;
}

/**
 * checkAuth is a small wrapper around authStore getAuthStorage and setAuthStorage
 */
export const checkAuth = (ctx: MiddlewareCtx) => {
  if (!ctx.authStore.isAuth) {
    const { user, access_token, remember } = ctx.authStore.getAuthStorage();
    if (user && access_token) {
      ctx.authStore.authUser = user;
      ctx.authStore.remember = remember;
      ctx.authStore.setAuthStorage(access_token);
    }
  }
};

/** auth redirects user to login page if they are not authenticated */
export const auth = (ctx: MiddlewareCtx) => {
  if (!ctx.authStore.isAuth) {
    ctx.next({ name: "login" });
    return false;
  }
  return true;
};

/** guets blocks authenticated user */
export const guest = (ctx: MiddlewareCtx) => {
  if (ctx.authStore.isAuth) {
    ctx.next(ctx.from);
    return false;
  }
  return true;
};
