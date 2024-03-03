<script setup lang="ts">
import { RequestResponse } from "@/utils/api";

defineProps({
  onFileChange: Function,
  onDelete: Function,
  rrDelete: RequestResponse,
  rrSubmit: { type: RequestResponse, required: true },
  onReset: Function,
});
</script>

<template>
  <div class="actions">
    <div class="field border small-margin prefix" v-if="onFileChange">
      <i>image</i>
      <input type="file" @change="onFileChange" accept=".jpeg,.jpg,.png" />
      <input type="text" id="image" name="image" placeholder="Change Image" />
    </div>
    <button
      v-if="rrDelete && onDelete"
      type="button"
      class="red small-opacity responsive"
      @click="onDelete(rrDelete)"
      :disabled="rrSubmit.loading || rrDelete.loading"
    >
      {{ rrDelete.loading ? "Deleting..." : "Delete" }}
      <i v-if="!rrDelete.loading">delete</i>
      <progress v-else class="circle small white-text"></progress>
    </button>
    <button
      type="button"
      class="inverted responsive"
      @click="onReset"
      v-if="onReset"
    >
      Reset
      <i>cancel</i>
    </button>
    <button
      type="submit"
      class="secondary responsive"
      :disabled="rrSubmit.loading || rrDelete?.loading"
    >
      {{ rrSubmit.loading ? "Processing..." : "Done" }}
      <i v-if="!rrSubmit.loading">done</i>
      <progress v-else class="circle small"></progress>
    </button>
  </div>
</template>

<style scoped>
button {
    margin: 0.5rem;
}

@media screen and (min-width: 992px) {
  .actions {
    position: fixed;
    width: 15rem;
  }
}
</style>