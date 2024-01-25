import { useApi } from "@/utils/api";
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

  /** fetchLatestPosts is a small wrapper around useApi and setLatestPosts */
  function fetchLatestPosts() {
    let data = ref<{posts: Post[]} | null>(null);
    let errors = ref(null);
    let status = ref(0); 
    
    if (latestPosts.value.length == 0) {
      ({data, errors, status} = useApi("/posts?order=desc&limit=5&total=5")); 
      setLatestPosts(data.value?.posts);
    }

    return { data, errors, status };
  }

  return {
    latestPosts,
    setLatestPosts,
    fetchLatestPosts,
  };
});
