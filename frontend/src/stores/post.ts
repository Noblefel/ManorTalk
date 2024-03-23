import { RequestResponse, Api } from "@/utils/api";
import { defineStore } from "pinia";
import { ref, type Ref } from "vue";
import { type User, useUserStore } from "./user";
import { Validator } from "@/utils/validator";
import { toast } from "@/utils/helper";
import { useRouter, type RouteLocation } from "vue-router";

export interface Post {
  id?: number;
  title: string;
  slug?: string;
  excerpt: string;
  image?: string;
  content: string;
  created_at?: string;
  updated_at?: string;
  category_id: number;
  category: Category;
  user_id: number;
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
  image: File | null;
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
  const latestPosts = ref([] as Post[]);
  const viewedPost = ref<Post | null>(null);
  const categories = ref([] as Category[]);
  const userStore = useUserStore();
  const router = useRouter();

  async function fetchHomePosts(rr: Ref<RequestResponse>) {
    rr.value.loading = true;
    try {
      const res = await Api.get("/posts?order=desc&limit=5&total=5");
      latestPosts.value = res.data.data.posts;
    } catch (e) {
      rr.value.handleErr(e);
    } finally {
      rr.value.loading = false;
    }
  }

  async function fetchPosts(rr: Ref<RequestResponse>, path = "") {
    if (path) path = path.slice(path.indexOf("?") + 1);
    rr.value.loading = true;
    try {
      const res = await Api.get("/posts?" + path);
      rr.value.data = res.data.data;
    } catch (e) {
      rr.value.handleErr(e);
    } finally {
      rr.value.loading = false;
    }
  }

  async function fetchPost(rr: Ref<RequestResponse>, to: RouteLocation) {
    if (viewedPost.value?.slug == to.params.slug) {
      return;
    }

    rr.value.loading = true;
    try {
      const res = await Api.get("/posts/" + to.params.slug);
      viewedPost.value = res.data.data;
    } catch (e) {
      rr.value.handleErr(e);
    } finally {
      rr.value.loading = false;
    }
  }

  async function fetchProfilePosts(
    rr: Ref<RequestResponse>,
    cursor: Ref<Number>
  ) {
    const id = userStore.viewedUser?.id;
    const params = `?cursor=${cursor.value}&user=${id}&limit=8&order=desc`;
    rr.value.loading = true;
    try {
      const res = await Api.get("/posts" + params);
      const next = res.data.data.posts;
      const prev = userStore.viewedUser?.posts;
      if (prev) {
        userStore.viewedUser!.posts = [...prev, ...next];
      } else {
        userStore.viewedUser!.posts = next;
      }

      cursor.value = next.slice(-1)[0].id || 0;
    } catch (e) {
      rr.value.handleErr(e);
    } finally {
      rr.value.loading = false;
    }
  }

  async function fetchCategories() {
    if (categories.value.length != 0) return;

    const rr = new RequestResponse();
    try {
      const res = await Api.get("/posts/categories");
      categories.value = res.data.data;
    } catch (e) {
      rr.handleErr(e);
    }
  }

  async function create(form: CreatePost, rr: RequestResponse) {
    const f = new Validator(form)
      .required("title", "content")
      .strMinLength("title", 10)
      .strMinLength("content", 50)
      .strMaxLength("title", 50);

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    }

    rr.loading = true;
    try {
      const res = await Api.post("/posts", form, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      toast(res.data.message, "green white-text");
      router.push({
        name: "blog.show",
        params: { slug: res.data.data.slug },
      });
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  async function update(form: CreatePost, rr: RequestResponse) {
    const f = new Validator(form)
      .required("title", "content")
      .strMinLength("title", 10)
      .strMinLength("content", 50)
      .strMaxLength("title", 50);

    if (!f.isValid()) {
      rr.errors = f.errors;
      toast("Some fields are invalid");
      return;
    }

    rr.loading = true;
    try {
      const res = await Api.patch("/posts/" + viewedPost.value?.slug, form, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      viewedPost.value = null;
      toast(res.data.message, "green white-text");
      router.push({
        name: "blog.show",
        params: { slug: res.data.data },
      });
    } catch (e) {
      rr.handleErr(e);
    } finally {
      rr.loading = false;
    }
  }

  async function deletePost(rr: RequestResponse) {
    if (!confirm("Are you sure? ")) return;

    rr.loading = true
    try {
      const res = await Api.delete("/posts/" + viewedPost.value?.slug) 
      viewedPost.value = null
      toast(res.data.message, "green white-text")
      router.push({name: "blog"})      
    } catch (e) {
      rr.handleErr(e)
    } finally {
      rr.loading = false
    }
  }

  return {
    latestPosts,
    viewedPost,
    categories,
    fetchPost,
    fetchPosts,
    fetchHomePosts,
    fetchProfilePosts,
    fetchCategories,
    create,
    update,
    deletePost,
  };
});
