import type { RequestResponse } from "@/utils/api";
import { defineStore } from "pinia";
import { ref, type Ref } from "vue";
import { type User, useUserStore } from "./user";
import { Validator } from "@/utils/validator";
import { toast } from "@/utils/helper";
import { useRouter } from "vue-router";

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
  id?: number;
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

export interface CreatePost {
  title: string;
  excerpt: string;
  content: string;
  category_id: number;
}

export const sampleContent = `### 🌟 Lorem Ipsum 👋

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.* 

###### 🌍 Fames Mauris

📚 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla commodo nisl sed odio hendrerit, sit amet dignissim libero fringilla. Fusce vehicula enim eget mauris suscipit, at porta nibh fermentum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed euismod turpis eget nisl molestie, id feugiat libero scelerisque. Curabitur convallis augue eu nisi fringilla, id vestibulum velit finibus. Pellentesque sit amet aliquet justo.

 
###### ☕ Nullam aliquam

🌱 Vivamus in libero varius, feugiat libero ut, pharetra lorem. Sed nec lacinia ante. Nam at mauris non libero bibendum lobortis. Nullam aliquam erat a tellus placerat, non ultricies libero ultrices. Integer ullamcorper, sem nec ultricies auctor, nisi mauris sollicitudin quam, eu eleifend nunc ipsum sed enim.   

**Code**
\`\`\` js
var foo = function (bar) {
  return bar++;
};
\`\`\`

___

###### 🎨 Curabitur convallis  

🎶 Ut at leo eu libero posuere eleifend. Cras congue vestibulum magna, id euismod nisl lacinia non. Donec ac dapibus lectus, eu congue enim. Integer non tellus ipsum. Sed malesuada sapien et odio fermentum, et ultricies purus efficitur.

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.* `;

export const usePostStore = defineStore("post", () => {
  /** new posts that are shown in the home page and sidebars */
  const latestPosts = ref([] as Post[]);
  const categories = ref([] as Category[]);
  const userStore = useUserStore();
  const postStore = usePostStore();
  const router = useRouter();

  /** fetchHomePosts will get newest posts and save it into the store */
  function fetchHomePosts(rr: RequestResponse) {
    if (latestPosts.value.length != 0) return;
    rr.useApi("get", "/posts?order=desc&limit=5&total=5").then(() => {
      if (rr.status != 200) return;

      const posts = (rr.data as unknown as { posts: Post[] }).posts;
      latestPosts.value = posts;
    });
  }

  /** fetchPosts will get posts with query parameters filter */
  function fetchPosts(rr: RequestResponse, path: string) {
    let query;
    if (path) query = path.slice(path.indexOf("?") + 1);
    return rr.useApi("get", "/posts?" + query);
  }

  /** fetchProfilePosts will get user's posts with cursor */
  function fetchProfilePosts(
    rr: RequestResponse,
    posts: Ref<Post[]>,
    cursor: Ref<Number>
  ) {
    let uId = userStore.viewedUser?.id;
    const params = `?cursor=${cursor.value}&user=${uId}&limit=9`;
    postStore.fetchPosts(rr, params).then(() => {
      const newPosts = (rr.data as any).posts as Post[];
      cursor.value = (newPosts.slice(-1)[0].id || 0) + 1;
      posts.value = [...posts.value, ...newPosts];
    });
  }

  /** fetchCategories will get all categories and save it into the store */
  function fetchCategories(rr: RequestResponse) {
    if (categories.value.length != 0) return;

    rr.useApi("get", "/posts/categories").then(() => {
      if (rr.status != 200) return;

      categories.value = rr.data as unknown as Category[];
    });
  }

  /** createPost will validates the form and stores the new post  */
  function createPost(form: CreatePost, rr: RequestResponse) {
    const f = new Validator(form)
      .required("title", "content")
      .strMinLength("title", 10)
      .strMinLength("content", 50)
      .strMaxLength("title", 255)
      .strMaxLength("title", 255);

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    } 

    rr.useApi("post", "/posts", form).then(() => {
      if (rr.status != 201) return;
      const slug = (rr.data as any as Post).slug;
      router.push({ name: "blog.show", params: { slug: slug } });
    });
  }

  return {
    latestPosts,
    categories,
    fetchHomePosts,
    fetchPosts,
    fetchProfilePosts,
    fetchCategories,
    createPost,
  };
});
