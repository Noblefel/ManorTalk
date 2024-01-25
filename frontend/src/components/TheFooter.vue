<template>
  <footer>
    <article class="center">
      <div class="text-wrapper">
        <p class="font-size-2 no-opacity font-700">
          <span class="orange-text">Subscribe </span>
          Newsletter 📰
        </p>
        <p>
          To be updated with all the latest news, trending discussions and
          special announcements.
        </p>
        <div class="field border round white prefix suffix">
          <i>mail</i>
          <input type="text" placeholder="Enter Your Email" />
          <i>send</i>
        </div>
      </div>
    </article>

    <div class="padding center-align">
      <a
        href="https://github.com/Noblefel/ManorTalk"
        class="button background"
        target="_blank"
      >
        <img src="@/assets/images/github.png" height="75px" alt="Github logo" />
      </a>
    </div>
  </footer>
</template>

<style scoped>
footer {
  background-color: black;
  min-height: 14rem;
}

article {
  background-color: #ffeec6;
  color: #ff9d48;
  width: 45rem;
  max-width: 100%;
  top: -7.5rem;
  margin-bottom: -7.5rem;
  .text-wrapper {
    max-width: 80%;
    margin: auto;
    text-align: center;

    & p {
      font-size: 1.2rem;
      font-weight: 600;
      color: black;
      opacity: 0.7;
    }
  }
}
</style>
