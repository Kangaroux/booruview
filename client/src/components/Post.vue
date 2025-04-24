<script setup lang="ts">
const { post } = defineProps<{
    post: Post;
}>();
</script>

<template>
    <img :src="post.preview_url" />
    <p>{{ post.tags }}</p>
</template>

<style scoped></style>
