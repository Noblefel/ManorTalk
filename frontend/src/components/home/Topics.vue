<template>
  <section id="home-topics">
    <div class="grid">
      <div class="s12 m12 l6">
        <div class="grid">
          <button class="s6 m6 l6 glow">
            <i>post</i>
            Blog Posts
          </button>
          <button class="s6 m6 l6 glow">
            <i>chat</i>
            Join Discussions
          </button>
          <button class="s6 m6 l6 glow">
            <i>light_mode</i>
            Placeholder
          </button>
          <button class="s6 m6 l6 glow">
            <i>dark_mode</i>
            Placeholder
          </button>
        </div>
      </div>
      <div class="s12 m12 l6 middle-align center-align">
        <h4 class="wrap font-size-title">
          Select which
          <span class="orange-text">topic</span>
          to dive into
        </h4>
      </div>
    </div>
  </section>
</template>

<style scoped>
#home-topics {
  padding: 3rem 1rem;

  .grid {
    margin: auto;
    max-width: 1100px;

    & button {
      font-size: 1.1rem;
      margin: 0 ;
      height: 7.5rem;
      display: flex;
      flex-direction: column; 
      background-size: cover;
      background-position: center;
      transition: 0.1s all ease-in; 
      filter: grayscale(0.5);

      &:hover {
        filter: grayscale(0);
      }

      &:nth-child(1) {
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
        url("@/assets/images/stock_1.jpg");
        background-size: cover;
      }

      &:nth-child(2) {
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
        url("@/assets/images/stock_2.jpg");
        background-size: cover;
      }

      &:nth-child(3) {
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
        url("@/assets/images/stock_3.jpg");
        background-size: cover;
      }

      &:nth-child(4) {
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
        url("@/assets/images/stock_1.jpg");
        background-size: cover;
      }
    }
  }

 
}
</style>
