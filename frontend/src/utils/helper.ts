import type { User } from "@/stores/user";
import { useRoute } from "vue-router";

/**
 *  Very small wrapper around BeerCSS ui() function .
 *
 * see docs https://github.com/beercss/beercss/blob/main/docs/DIALOG.md#method-4
 */
export const beerUi = (id: string): void => {
  ui("#" + id);
};

/**  activeRoute returns "active" if the current route name match  .*/
export const activeRoute = (routeName: string): string => {
  const route = useRoute();

  if (route.name == routeName) {
    return "active";
  }

  return "";
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

  if (user.avatar) return user.avatar;

  const name = (user.name ?? user.username).split(/[\s_\-]/).join("+");

  return `https://ui-avatars.com/api/?name=${name}&background=random&size=120&color=fff`;
};
