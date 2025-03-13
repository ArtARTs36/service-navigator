<template>
  <div class="bg-light rounded">
    <p style="padding-top: 10px">
      <span class="h4">Images ({{ images.length }})</span>
      <span class="h5">
        <a href="/images/refresh"><i class="fa-solid fa-arrows-rotate link-action-refresh "></i></a>
      </span>
    </p>

    <ImagesList :images="images"></ImagesList>
  </div>
</template>

<script>
import ImagesList from "@/components/pages/images/List.vue";

export default {
  name: "ImagesPageIndex",
  components: {ImagesList},

  data() {
    return {
      images: [
        {
          id: "1",
          name: "Redis",
          size: {
            readable: "2 GB",
          },
          languages: [
            {
              name: "Go",
              version: "1.23",
            },
          ],
          vcs: {
            url: "https://github.com/redis/redis",
            type: "github",
            host: "github.com",
          },
          registry: {
            url: "https://hub.docker.com/_/redis",
          },
        }
      ],
    }
  },
}
</script>
