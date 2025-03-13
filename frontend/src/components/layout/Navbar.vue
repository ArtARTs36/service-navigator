<template>
  <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
    <li v-for="link in config.links">
      <a :href="link.url" class="nav-link px-2 navbar-link">{{ link.title }}</a>
    </li>
  </ul>

  <NavbarProfile :username="username" :config="config.profile"></NavbarProfile>
</template>

<script>
import NavbarProfile from "@/components/layout/NavbarProfile.vue";

export default {
  name: "Navbar",

  components: {
    NavbarProfile,
  },

  props: {
    username: {
      type: String,
    },

    config: {
      type: Object,
    },
  },
}
</script>
