<script setup lang="ts">
import AuthCard from "@/components/auth/AuthCard.vue";
import { ref } from "vue";
import { RouterLink } from "vue-router";
import { RequestResponse } from "@/utils/api";
import { useAuthStore, type RegisterForm } from "@/stores/auth";

const authStore = useAuthStore();
const showPassword = ref(false);

const form = ref<RegisterForm>({
  email: "",
  password: "",
  password_repeat: "",
});

const rr = ref(new RequestResponse());
</script>

<template>
  <AuthCard title="Create new account 🎉">
    <form @submit.prevent="authStore.register(form, rr)"> 
      <div class="padding">
        <label for="email" class="font-size-0-9 font-600">Email</label>
        <div class="field border no-margin prefix">
          <i>mail</i>
          <input
            type="text"
            name="email"
            id="email"
            autocomplete="off"
            placeholder="test@example.com"
            v-model.trim="form.email"
          />
          <span class="error" v-if="rr.errors?.email">
            {{ rr.errors.email[0] }}
          </span>
        </div>

        <div class="space"></div>

        <label for="password" class="font-size-0-9 font-600">Password</label>
        <div class="field border no-margin prefix suffix">
          <i>key</i>
          <input
            :type="showPassword ? 'text' : 'password'"
            name="password"
            id="password" 
            placeholder="password"
            v-model.trim="form.password"
          />
          <span class="error" v-if="rr.errors?.password">
            {{ rr.errors.password[0] }}
          </span>
          <i @click="showPassword = !showPassword" class="cursor-pointer z-2">
            visibility{{ showPassword == true ? "_off" : "" }}
          </i>
        </div>

        <div class="space"></div>

        <label for="password" class="font-size-0-9 font-600">Repeat Password</label>
        <div class="field border no-margin prefix suffix">
          <i>key</i>
          <input
            type="password"
            name="password_repeat"
            id="password_repeat" 
            v-model.trim="form.password_repeat"
          />
          <span class="error" v-if="rr.errors?.password_repeat">
            {{ rr.errors.password_repeat[0] }}
          </span> 
        </div>

        <div class="space"></div> 
      </div>
 
      <button class="responsive" :disabled="rr.loading">
        {{ rr.loading ? "Registering..." : "Register" }}
        <i v-if="!rr.loading">person_add</i>
        <progress v-else class="circle white-text small"></progress>
      </button>

      <p class="center-align font-600 font-size-0-9">
        Already have an account?
        <RouterLink :to="{ name: 'login' }" class="orange-text">
          Login
        </RouterLink>
      </p>
    </form>
  </AuthCard>
</template> 