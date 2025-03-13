<template>
<Header></Header>

<main>
  <div class="container">
    <router-view />
  </div>
</main>
</template>

<script>
import Header from "@/components/layout/Header.vue";

export default {
  name: "App",

  components: {
    Header,
  }
}
</script>
