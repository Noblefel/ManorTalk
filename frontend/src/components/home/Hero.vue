<template>
  <section id="home-hero">
    <div class="grid">
      <div class="s12 s m center-align">
        <img src="@/assets/images/hero.png" alt="Hero Vector" />
      </div>
      <div class="s12 m12 l6">
        <div class="text-wrapper middle">
          <h1>Welcome</h1>
          <p class="font-size-1-25 font-500">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi in
            sapien luctus libero ultricies egestas vel vitae ligula. Aliquam
            posuere condimentum ex, quis luctus libero interdum eget
          </p>
          <div class="space"></div>
          <button class="large no-margin">Get Started&nbsp; ✏️</button>
        </div>
      </div>
      <div class="l6 l">
        <img src="@/assets/images/hero.png" alt="Hero Vector" />
      </div>
    </div>
  </section>
</template>

<style scoped>
#home-hero {
  display: flex;
  justify-content: center;
  align-items: center;

  & .grid {
    width: 100%;
    max-width: 1000px;
  }

  & img {
    width: clamp(15rem, calc(13rem + 17vw), 40rem);
    max-width: 95%;
  }

  & .text-wrapper {
    max-width: 750px;
    margin: auto;
    padding-bottom: 3rem;
    padding-left: clamp(0.5rem, calc(0.5rem + 1.5vw), 3rem);
  }
}
</style>
