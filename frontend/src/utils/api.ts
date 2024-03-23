import type { AxiosInstance } from "axios";
import axios from "axios";
import type { FormErrors } from "./validator";
import { toast } from "./helper";
import { useAuthStore } from "@/stores/auth";

export const Api: AxiosInstance = axios.create({
  baseURL: "http://localhost:8080/api",
  withCredentials: true,
});

Api.interceptors.request.use(async (config) => {
  const token =
    localStorage.getItem("access_token") ??
    sessionStorage.getItem("access_token");
  config.headers.Authorization = token;
  return config;
});

Api.interceptors.response.use(
  (res) => {
    return res;
  },
  async (e) => {
    const originalReq = e.config;
    const expired =
      e.response?.status == 401 && e.response?.data.message == "Token Expired";

    if (expired && !originalReq._retry) {
      originalReq._retry = true;
      const authStore = useAuthStore();
      try {
        const res = await Api.post("/auth/refresh");
        authStore.setAuthStorage(res.data.data.access_token);

        return Api(originalReq);
      } catch (errRefresh) {
        authStore.reset();
        window.location.reload();

        throw errRefresh;
      }
    }

    return Promise.reject(e);
  }
);

export class RequestResponse {
  data: null;
  errors: FormErrors | null;
  message: string | null;
  status: number;
  loading: boolean;

  constructor() {
    this.data = null;
    this.errors = null;
    this.message = null;
    this.status = 0;
    this.loading = false;
  }

  public handleErr(e: any, showToast = true) {
    console.log(e);
    if (e.response) {
      this.message = e.response.data.message;
      this.status = e.response.status;
      this.errors = e.response.data.errors ?? 1;
    } else if (e.request) {
      this.message = e.message;
      this.errors = e;
    } else {
      this.message = "Unexpected Error";
      this.errors = e;
    }
  
    if (showToast && this.message && this.message != "Token Expired") {
      toast(this.message);
    }
  } 
}
