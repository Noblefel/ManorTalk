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
export const activeRoute = (routeName: string) : string => {
  const route = useRoute();

  if (route.name == routeName) {
    return "active";
  }

  return ""
};

/** toast will create a pop up div element with the provided message
 *  that last 7 seconds
 */
export const toast = (messages: string, className: string = "error") => {
    const id = `toast-${Math.ceil(Math.random() * 1000)}` 

    let toast = document.createElement("div")
    toast.setAttribute("class", `snackbar ${className}`)
    toast.setAttribute("id", id)
    toast.innerHTML = messages

    document.body.appendChild(toast)

    ui('#' + id, 7000) 

    setTimeout(() => {
        toast.remove()
    }, 7000);
}