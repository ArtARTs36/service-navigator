<template>
  <header class="p-3 border-bottom" style="background: #fff">
    <div class="container">
      <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
        <a href="/" class="d-flex align-items-center mb-2 mb-lg-0 text-dark text-decoration-none logo">
          {{ config.appName }}
        </a>

        <Navbar :username="username" :config="config.navbar"></Navbar>
      </div>
    </div>
  </header>
</template>

<script>
import Navbar from "@/components/layout/Navbar.vue";

export default {
  components: {
    Navbar,
  },

  data() {
    return {
      username: "developer",

      config: {
        appName: "Service Navigator",
        navbar: {
          links: [
            {
              title: "Services",
              url: "/",
            },
            {
              title: "Images",
              url: "/images",
            },
            {
              title: "Volumes",
              url: "/volumes",
            },
            {
              title: "Github",
              url: "https://github.com/artarts36/service-navigator",
            },
          ],

          search: {
            providers: [
              {

              }
            ],
          },

          profile: {
            links: [],
          },
        },
      },
    }
  },
}
</script>
