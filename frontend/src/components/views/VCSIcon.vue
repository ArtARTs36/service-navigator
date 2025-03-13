<template>
  <i class="fa-brands" :class="iconClass"></i>
</template>

<script>
export default {
  name: "VCSIcon",

  props: {
    type: {
      type: String,
    },
  },

  computed: {
    iconClass: function () {
      return `fa-${this.type}`
    }
  },
}
</script>
