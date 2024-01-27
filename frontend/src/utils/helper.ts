import { useRoute } from "vue-router";

/**
 *  Very small wrapper around BeerCSS ui() function .
 *
 * see docs https://github.com/beercss/beercss/blob/main/docs/DIALOG.md#method-4
 */
export const toggleUi = (id: string): void => {
  ui("#" + id);
};

/**  activeRoute returns "active" if the current route name match  .*/
export const activeRoute = (routeName: string) : string => {
  const route = useRoute();

  if (route.name == routeName) {
    return "active";
  }

  return ""
};
