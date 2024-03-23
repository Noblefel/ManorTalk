<script setup lang="ts">
import AuthCard from "@/components/auth/AuthCard.vue";
import { ref } from "vue";
import { RouterLink } from "vue-router";
import { RequestResponse } from "@/utils/api";
import { type LoginForm, useAuthStore } from "@/stores/auth";

const as = useAuthStore();
const showPassword = ref(false);

const form = ref<LoginForm>({
  email: "test@example.com",
  password: "password",
  remember_me: false,
});

const rr = ref(new RequestResponse());
</script>

<template>
  <AuthCard title="Sign in to continue 📝">
    <form @submit.prevent="as.login(form, rr)">
      <div class="space"></div>

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
            autocomplete="off"
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

        <div class="row">
          <label class="checkbox">
            <input
              type="checkbox"
              :checked="form.remember_me"
              v-model="form.remember_me"
            />
            <span class="font-500">Remember Me</span>
          </label>
          <div class="max right-align">
            <p class="orange-text font-600">Forgot Password</p>
          </div>
        </div>
      </div>

      <div class="space"></div>

      <button class="responsive" :disabled="rr.loading">
        {{ rr.loading ? "Logging in..." : "Login" }}
        <i v-if="!rr.loading">login</i>
        <progress v-else class="circle white-text small"></progress>
      </button>

      <p class="center-align font-600">
        New?
        <RouterLink :to="{ name: 'register' }" class="orange-text">
          Create an account
        </RouterLink>
      </p>
    </form>
  </AuthCard>
</template>

<style scoped>
p {
  font-size: 0.9rem;
}
</style>
