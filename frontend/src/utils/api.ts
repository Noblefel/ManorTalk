import type { AxiosInstance } from "axios";
import axios from "axios";
import { ref, type Ref } from "vue";

const Api: AxiosInstance = axios.create({
  baseURL: "http://localhost:8080/api",
  //   withCredentials: true,
});
 
export function useApi(url: string) {
  const data = ref(null);
  const errors = ref(null);
  const message = ref(null);
  const status = ref(0);

  const execute = async () => {
    try { 
      const res = await Api.get(url);
      data.value = res.data.data;
      message.value = res.data.message;
      status.value = res.status;
    } catch (e: any) {
      console.log(e);
      if (e.response) {
        message.value = e.response.data.message;
        status.value = e.response.status;
        errors.value = e.response.data.errors;
      } else {
        errors.value = e;
      }
    }
  };

  execute()

  return {
    data,
    errors,
    message,
    status,
  };
}
