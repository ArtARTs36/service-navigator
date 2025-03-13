<template>
  <div class="table-responsive">
    <table class="table bordered">
      <thead>
      <tr>
        <td scope="col">
          <strong>Name</strong>
        </td>
        <td scope="col">
          <strong>Size</strong>
        </td>
        <td scope="col">
          <strong>Actions</strong>
        </td>
        <td scope="col">
          <strong>Counters</strong>
        </td>
        <td>
          Language
        </td>
        <td scope="col">
          <strong>Links</strong>
        </td>
      </tr>
      </thead>

      <tbody>
      <ListItem v-for="image in images" :image="image"></ListItem>
      </tbody>
    </table>
  </div>
</template>

<script>
import ListItem from "@/components/pages/images/ListItem.vue";

export default {
  name: "ImagesList",
  components: {ListItem},

  props: {
    images: {
      type: Array,
    },
  },
}
</script>
