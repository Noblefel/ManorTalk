import type { User } from "@/stores/user";
import { useRoute, type Router } from "vue-router";

/**  activeRoute returns "active" if the current route name match  .*/
export const activeRoute = (routeName: string): string => {
  const route = useRoute();

  return route.name?.toString().startsWith(routeName) ? "active" : "";
};

/** toast will create a pop up div element with the provided message
 *  that last 7 seconds
 */
export const toast = (messages: string, className: string = "error") => {
  const id = `toast-${Math.ceil(Math.random() * 1000)}`;

  let toast = document.createElement("div");
  toast.setAttribute("class", `snackbar ${className}`);
  toast.setAttribute("id", id);
  toast.innerHTML = messages;

  document.body.appendChild(toast);

  ui("#" + id, 7000);

  setTimeout(() => {
    toast.remove();
  }, 7000);
};

/** getAvatar will return image url based on user's name/username.
 *  Will skip if user already has an avatar
 */
export const getAvatar = (user: User | null) => {
  if (!user)
    return `https://ui-avatars.com/api/?name=guest&background=ffeec6&size=120&color=ff9d48&bold=true`;

  if (user.avatar) return "http://localhost:8080/images/avatar/" + user.avatar;

  const name = (user.name ?? user.username).split(/[\s_\-]/).join("+");

  return `https://ui-avatars.com/api/?name=${name}&background=random&size=120&color=fff`;
};

/** changeParam will set new query parameter and modify the url */
export const changeParam = (router: Router, key: string, value: any) => {
  let params = new URLSearchParams(window.location.search);
  if (["search", "limit"].includes(key)) params.delete("page");
  if (key == "page") window.scrollTo(0, 150);
  params.set(key, value);
  router.replace({ query: Object.fromEntries(params) });
};

/**  ============ NEED IMPROVEMENTS ============
 *
 * getPageNumbers will return array of numbers from before and after the current page
 */
export const getPageNumbers = (current: number, last: number) => {
  let beforeNum = Math.max(current - 3, 1);
  let afterNum = Math.min(current + 3, last);

  let additionalNum = {
    before: 3 + (current - afterNum),
    after: 3 - (current - beforeNum),
  };

  let before = [];
  for (let i = beforeNum - additionalNum.before; i < current; i++) {
    if (i <= 0) continue;
    before.push(i);
  }

  let after = [];
  for (let i = current + 1; i <= afterNum + additionalNum.after; i++) {
    if (i > last) break;
    after.push(i);
  }

  return { before, after };
};
