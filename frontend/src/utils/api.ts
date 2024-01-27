import type { AxiosInstance } from "axios";
import axios from "axios";
import type { FormErrors } from "./validator";

const Api: AxiosInstance = axios.create({
  baseURL: "http://localhost:8080/api",
  //   withCredentials: true,
});

/** RequestResponse is a utility class for handling states while using the  API */
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

  async useApi(method: string, url: string, body: any = null) {
    this.loading = true;

    let req;
    switch (method) {
      case "get":
        req = Api.get(url);
        break;
      case "post":
        req = Api.post(url, body);
        break;
      default:
        req = Api.get(url);
        break;
    }

    return await req
      .then((res) => {
        this.data = res.data.data;
        this.message = res.data.message;
        this.status = res.status;
      })
      .catch((e) => {
        console.log(e);
        if (e.response) {
          this.message = e.response.data.message;
          this.status = e.response.status;
          this.errors = e.response.data.errors ?? 1;
        } else {
          this.errors = e;
        }
      })
      .finally(() => {
        this.loading = false;
      });
  }
}
