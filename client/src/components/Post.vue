<script setup lang="ts">
const { post } = defineProps<{
    post: Post;
}>();
</script>

<template>
    <img
        :src="post.lowres_url"
        :width="post.low_res_width"
        :height="post.low_res_height"
        loading="lazy"
    />
    <p>{{ post.tags }}</p>
</template>

<style scoped></style>
