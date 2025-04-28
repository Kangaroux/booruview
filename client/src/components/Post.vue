<script setup lang="ts">
import { computed, defineProps } from "vue";

const { post } = defineProps<{
    post: Post;
}>();

const image = computed<{url: string, width: number, height: number}>(() => {
    if (post.lowres_url.length > 0) {
        return {
            url: post.lowres_url,
            width: post.lowres_width,
            height: post.lowres_height,
        };
    }

    return {
        url: post.image_url,
        width: post.width,
        height: post.height,
    };
});
</script>

<template>
    <img
        :src="image.url"
        :width="image.width"
        :height="image.height"
        loading="lazy"
        referrerpolicy="no-referrer"
    />
    <p>{{ post.tags }}</p>
</template>

<style scoped></style>
