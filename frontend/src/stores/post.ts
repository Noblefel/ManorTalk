// import { useApi } from "@/utils/api";
import type { RequestResponse } from "@/utils/api";
import { defineStore } from "pinia";
import { ref } from "vue";

export interface Post {
  id?: number;
  title: string;
  slug?: string;
  excerpt: string;
  content: string;
  category: { name: string };
  created_at?: string;
  updated_at?: string;
}

export const usePostStore = defineStore("post", () => {
  /** new posts that are shown in the home page and sidebars */
  const latestPosts = ref([] as Post[]);

  function setLatestPosts(posts: Post[] | undefined) {
    if (posts == undefined) {
      return;
    }
    latestPosts.value = posts;
  }

  /** fetchLatestPosts is a small wrapper around rr.useApi and setLatestPosts */
  function fetchLatestPosts(rr: RequestResponse) {
    if (latestPosts.value.length == 0) {
      rr.useApi("get", "/posts?order=desc&limit=5&total=5").then(() => {
        const posts = (rr.data as null | { posts: Post[] })?.posts;
        setLatestPosts(posts);
      });
    }
  }

  return {
    latestPosts,
    setLatestPosts,
    fetchLatestPosts,
  };
});
