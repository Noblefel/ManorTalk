// import { useApi } from "@/utils/api";
import type { RequestResponse } from "@/utils/api";
import { defineStore } from "pinia";
import { ref } from "vue";
import type { User } from "./user";

export interface Post {
  id?: number;
  title: string;
  slug?: string;
  excerpt: string;
  content: string;
  created_at?: string;
  updated_at?: string;
  category: Category;
  user: User;
}

export interface Category {
  name: string;
  slug: string;
}

export interface PaginationMeta {
  current_page: number;
  total: number;
  last_page: number;
  limit: number;
  offset: number;
}

export const usePostStore = defineStore("post", () => {
  /** new posts that are shown in the home page and sidebars */
  const latestPosts = ref([] as Post[]);
  const categories = ref([] as Category[]);

  /** fetchHomePosts will get newest posts and save it into the store */
  function fetchHomePosts(rr: RequestResponse) {
    if (latestPosts.value.length == 0) {
      rr.useApi("get", "/posts?order=desc&limit=5&total=5").then(() => {
        if (rr.status != 200) return;

        const posts = (rr.data as unknown as { posts: Post[] }).posts;
        latestPosts.value = posts;
      });
    }
  }

  /** fetchPosts will get posts with query parameters filter */
  function fetchPosts(rr: RequestResponse, path: string) {
    let query;
    if (path) query = path.slice(path.indexOf("?") + 1);
    rr.useApi("get", "/posts?" + query);
  }

  /** fetchCategories will get all categories and save it into the store */
  function fetchCategories(rr: RequestResponse) {
    if (categories.value.length == 0) {
      rr.useApi("get", "/posts/categories").then(() => {
        if (rr.status != 200) return;

        categories.value = rr.data as unknown as Category[];
      });
    }
  }

  return {
    latestPosts,
    categories,
    fetchHomePosts,
    fetchPosts,
    fetchCategories,
  };
});
