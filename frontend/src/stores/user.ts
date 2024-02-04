import type { RequestResponse } from "@/utils/api";
import { toast } from "@/utils/helper";
import { Validator } from "@/utils/validator";
import { defineStore } from "pinia";
import { computed } from "vue";

export interface User {
  id: number;
  name: string;
  username: string;
  avatar: string;
  email: string;
  password: string;
  created_at?: string;
  updated_at?: string;
}

export const useUserStore = defineStore("user", () => {
  /** checkUsername validates the username and send request to check its availability */
  function checkUsername(username: string, rr: RequestResponse) {
    const f = new Validator({ username: username })
      .required("username")
      .strMinLength("username", 3)
      .strMaxLength("username", 40);

    if (!f.isValid()) {
      rr.errors = f.errors;
      return;
    }

    rr.useApi("post", "/user/check-username", f.form).then(() => {
      if (rr.status != 200) return;
      rr.errors = null;

      if (rr.message) toast(rr.message, "green white-text");
    });
  }

  return {
    checkUsername,
  };
});
