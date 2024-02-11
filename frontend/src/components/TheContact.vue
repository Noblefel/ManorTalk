<template>
  <section id="home-contact">
    <div class="wrapper">
      <h4 class="font-size-title wrap center-align">
        <span class="orange-text">Contact</span>
        Us
      </h4>

      <div class="row scroll">
        <div>
          <i class="secondary">mail</i>
          <p>test@example.com</p>
        </div>
        <div>
          <i class="secondary">phone</i>
          <p>111-222222-33333</p>
        </div>
        <div>
          <i class="secondary">location_on</i>
          <p>Addr X, ManorTalk, 123</p>
        </div>
      </div>

      <div class="grid middle-align">
        <div class="s12 m12 l6 center-align">
          <img src="@/assets/images/logo.png" alt="House Vector" />
        </div>
        <div class="s12 m12 l6"> 
          <p class="font-600 font-size-1-25">Get in touch</p>
          <div class="space"></div>
          <div class="grid">
            <div class="field border prefix s12 m6 l6 no-margin">
              <i>account_circle</i>
              <input type="text" name="contact-name" placeholder="Name" />
            </div>
            <div class="field border prefix s12 m6 l6 no-margin">
              <i>mail</i>
              <input type="text" name="contact-email" placeholder="Email" />
            </div>
          </div>
          <div class="space"></div>
          <div class="field border">
            <input type="text" name="contact-subject" placeholder="Subject" />
          </div>
          <div class="field textarea border large">
            <textarea placeholder="Message"></textarea>
          </div>
          <div class="space"></div>
          <button class="no-margin">Send Message</button>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.row.scroll {
  gap: 2rem;
  margin: 2.5rem 0;
  justify-content: center;
  text-align: center;
  font-weight: 600;

  i {
    padding: 0.75rem;
    border-radius: 10px;
    background-color: var(--primary);
  }
}

@media screen and (max-width: 600px) {
  .row.scroll {
    justify-content: start;
  }
}

img {
  max-width: 90%;
  width: clamp(13rem, calc(8rem + 30vw), 30rem);
  opacity: 0.8;
}
</style>
