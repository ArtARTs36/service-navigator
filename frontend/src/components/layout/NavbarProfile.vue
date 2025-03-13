<template>
  <div class="dropdown text-end">
    <a href="#" class="d-block link-dark text-decoration-none dropdown-toggle" id="dropdownUser1" data-bs-toggle="dropdown" aria-expanded="false">
      {{ username }}
    </a>
    <ul class="dropdown-menu text-small" aria-labelledby="dropdownUser1">
      <li v-for="link in config.links">
        <a class="dropdown-item" :href="link.url">{{ link.title }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "NavbarProfile",

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
