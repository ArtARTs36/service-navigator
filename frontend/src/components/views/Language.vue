<template>
{{ language.name }}<span v-if="language.version">{{ language.version }}</span>
</template>

<script>
export default {
  name: "Language",

  props: {
    language: {
      type: Object,
    },
  },
}
</script>
