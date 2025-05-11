<script setup lang="ts">
import store from "@/store";
import { computed, ref } from "vue";

const { post } = defineProps<{
    post: Post;
}>();

const fetchingTags = ref(false);
const content = computed<{ url: string; width: number; height: number }>(() => {
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

const isVideo = computed(() => {
    return (
        content.value.url.endsWith(".mp4") ||
        content.value.url.endsWith(".webm")
    );
});
</script>

<template>
    <div class="post">
        <video
            class="content"
            :poster="post.thumbnail_url || post.lowres_url"
            :width="content.width"
            :height="content.height"
            preload="none"
            controls
            v-if="isVideo"
        >
            <source
                :src="content.url"
                type="video/mp4"
                v-if="content.url.endsWith('.mp4')"
            />
            <source
                :src="content.url"
                type="video/webm"
                v-if="content.url.endsWith('.webm')"
            />
        </video>

        <img
            class="content"
            :src="content.url"
            :width="content.width"
            :height="content.height"
            loading="lazy"
            @click="store.fullscreenPost = post"
            v-if="!isVideo"
        />
    </div>
</template>

<style scoped>
.post {
    break-inside: avoid;
}

.content {
    /* placeholder color */
    background-color: #444;
    width: 100%;
    height: auto;
}
</style>
