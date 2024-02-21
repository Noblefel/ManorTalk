<script setup lang="ts">
// @ts-ignore
import markdownit from "markdown-it";

defineProps({
  content: String,
});

const md = markdownit({breaks: true});
const render = md.render(`

### 🌟 Lorem Ipsum 👋

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.* 

###### 🌍 Fames Mauris

📚 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla commodo nisl sed odio hendrerit, sit amet dignissim libero fringilla. Fusce vehicula enim eget mauris suscipit, at porta nibh fermentum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed euismod turpis eget nisl molestie, id feugiat libero scelerisque. Curabitur convallis augue eu nisi fringilla, id vestibulum velit finibus. Pellentesque sit amet aliquet justo.

 
###### ☕ Nullam aliquam

🌱 Vivamus in libero varius, feugiat libero ut, pharetra lorem. Sed nec lacinia ante. Nam at mauris non libero bibendum lobortis. Nullam aliquam erat a tellus placerat, non ultricies libero ultrices. Integer ullamcorper, sem nec ultricies auctor, nisi mauris sollicitudin quam, eu eleifend nunc ipsum sed enim.   

**Code**
\`\`\` js
var foo = function (bar) {
  return bar++;
};
\`\`\`

___

###### 🎨 Curabitur convallis  

🎶 Ut at leo eu libero posuere eleifend. Cras congue vestibulum magna, id euismod nisl lacinia non. Donec ac dapibus lectus, eu congue enim. Integer non tellus ipsum. Sed malesuada sapien et odio fermentum, et ultricies purus efficitur.

*Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla posuere neque id magna pretium rutrum.* 
`);
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
    
    h1, h2, h3, h4, h5, h6 {
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
