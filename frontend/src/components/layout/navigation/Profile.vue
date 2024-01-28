<script setup lang="ts">
import { beerUi } from "@/utils/helper";
</script>

<template>
  <div class="nav-profile">
    <div class="row responsive">
      <img
        src="@/assets/images/default_pfp_240.png"
        alt=""
        width="50px"
        height="50px"
      />
      <div class="max right-align row">
        <button class="no-margin circle secondary">
          <i class="fill">manage_accounts</i>
          <div class="tooltip bottom">Edit Profile</div>
        </button>
        <button
          class="no-margin circle m s"
          @click="beerUi('nav-mobile-menu')"
        >
          <i>close</i>
        </button>
      </div>
    </div>
    <div class="">
      <p class="no-margin">A Sample Name</p>
      <p class="no-margin font-size-0-9 small-opacity">#a-sample-username</p>
    </div>
  </div>
</template>

<style scoped>
.nav-profile {
  margin-bottom: 0.5rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  font-weight: 600;
}
</style>
