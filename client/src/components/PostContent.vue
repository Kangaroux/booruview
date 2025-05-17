<script setup lang="ts">
import store from "@/store";
import { computed } from "vue";

const { cropped, maxHeight, post } = defineProps<{
    cropped: boolean;
    maxHeight: number;
    post: Post;
}>();

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
    <div class="post" :style="{ maxHeight: maxHeight + 'px' }">
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

        <div class="img-container" @click="store.fullscreenPost = post">
            <span
                class="crop-icon"
                title="post is cropped (too tall)"
                v-if="cropped"
            >
                <i class="bi bi-crop"></i>
            </span>
            <img
                class="content"
                :src="content.url"
                :width="content.width"
                :height="content.height"
                loading="lazy"
                v-if="!isVideo"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.post {
    break-inside: avoid;
    font-size: 0;
    position: relative;
}

.content {
    /* placeholder color */
    background-color: #444;
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.img-container {
    position: relative;
    width: 100%;
    height: 100%;
    cursor: pointer;
}

.crop-icon {
    font-size: 24px;
    position: absolute;
    right: 10px;
    top: 10px;
    background-color: black;
    opacity: 0.2;
    border-radius: 50%;
    padding: 0.6em;
    line-height: 0;

    .bi {
        display: block;
        line-height: 0;
        position: relative;
        left: -0.06em;
        top: -0.06em;
    }
}
</style>
