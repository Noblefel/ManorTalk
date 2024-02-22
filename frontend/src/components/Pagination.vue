<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, type PropType } from "vue";
import type { PaginationMeta } from "@/stores/post";
import { changeParam, getPageNumbers } from "@/utils/helper";

const router = useRouter();

const props = defineProps({
  paginationMeta: {
    type: Object as PropType<PaginationMeta>,
    required: true,
  },
});

const goToPage = ref(null);
const current = props.paginationMeta.current_page;
const last = props.paginationMeta.last_page;

const { before, after } = getPageNumbers(current, last);
</script>

<template>
  <div>
    <div class="row">
      <div class="m l">
        <a @click="changeParam(router, 'page', 1)"> &lt;&lt; </a>
      </div>
      <div class="row">
        <a v-for="page in before" @click="changeParam(router, 'page', page)">
          {{ page }}
        </a>

        <a class="primary">{{ current }}</a>

        <a v-for="page in after" @click="changeParam(router, 'page', page)">
          {{ page }}
        </a>
      </div>
      <div class="m l">
        <a @click="changeParam(router, 'page', last)"> {{ last }} </a>
      </div>
    </div>

    <div class="row center-align">
      <div class="field border go-to">
        <input
          type="number"
          class="no-padding center-align"
          id="go-to-page"
          v-model="goToPage"
        />
      </div>
      <button
        class="circle small-round secondary"
        @click="changeParam(router, 'page', goToPage)"
      >
        Go
      </button>
    </div>

    <div class="row center-align s">
      <a class="button small secondary" @click="changeParam(router, 'page', 1)">
        &lt;&lt;
      </a>
      <a
        class="button small secondary"
        @click="changeParam(router, 'page', last)"
      >
        >>
      </a>
    </div>
  </div>
</template>

<style scoped>
.row:first-of-type {
  justify-content: center;
  overflow: auto;
  flex-wrap: wrap;

  .row {
    gap: 0.5rem;
    overflow: auto;
  }
}

a {
  padding: 0.5rem 1rem;
  border-radius: 8px;
  transition: 0.2s all ease-in;

  &:hover {
    background-color: var(--secondary);
    color: var(--on-secondary);
  }
}

.go-to {
  width: 5rem;
  height: 2.5rem;
}
</style>
