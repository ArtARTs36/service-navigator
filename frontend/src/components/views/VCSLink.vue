<template>
  <a :href="url" :class="linkClass">
    <VCSIcon v-if="type" :type="type"></VCSIcon>
    <span v-else>{{ host }}</span>
  </a>
</template>

<script>
import VCSIcon from "@/components/views/VCSIcon.vue";

export default {
  name: "VCSLink",
  components: {VCSIcon},

  props: {
    type: {
      type: String,
    },
    host: {
      type: String,
    },
    url: {
      type: String,
    },
  },

  computed: {
    linkClass: function () {
      return `link-${this.type}`
    },
  },
}
</script>
