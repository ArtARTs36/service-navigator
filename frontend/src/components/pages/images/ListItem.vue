<template>
  <tr class="table-row">
    <td>
      {{ image.name }}
    </td>
    <td>
      {{ image.size.readable }}
    </td>
    <td>
      <a href="/images/remove?imageId={{ image.id }}"><i class="fa-solid fa-trash-can link-action-stop"></i></a>
    </td>
    <td>
      <Language v-for="lang in image.languages" :language="lang"></Language>
    </td>
    <td>
      no counters
    </td>
    <td>
      <VCSLink v-if="image.vcs" :url="image.vcs.url" :type="image.vcs.type" :host="image.vcs.host"></VCSLink>

      <span v-if="image.vcs">&nbsp;</span>

      <a v-if="image.registry && image.registry.url" :href="image.registry.url" class="link-docker">
        <i class="fa-brands fa-docker"></i>
      </a>
    </td>
  </tr>
</template>

<script>
import Language from "@/components/views/Language.vue";
import VCSLink from "@/components/views/VCSLink.vue";

export default {
  components: {
    Language,
    VCSLink,
  },

  props: {
    image: {
      type: Object,
    },
  },
}
</script>
