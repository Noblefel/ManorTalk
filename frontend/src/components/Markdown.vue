<script setup lang="ts">
// @ts-ignore
import markdownit from "markdown-it";

const props = defineProps({
  content: {
    type: String,
    requried: true,
  },
});

const md = markdownit({ breaks: true });
const render = md.render(props.content);
</script>

<template>
  <div class="markdown-body" v-html="render"></div>
</template>

<style>
.markdown-body {
  background-color: var(--background);
  color: var(--on-surface);
  font-size: 1.2rem;
  font-family: DMSans;

  pre {
    padding: 1rem;
    background-color: rgba(133, 133, 133, 0.068);
    border-radius: 8px;
    font-size: 1rem;
  }

  hr {
    opacity: 0.5;
    margin: 1rem 0;
  }

  p {
    margin: 1rem 0;
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    margin: 2rem 0;
  }
}

body.dark {
  .markdown-body {
    pre {
      color: var(--on-surface);
      border-radius: 8px;
    }
  }
}
</style>
