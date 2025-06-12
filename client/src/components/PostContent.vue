<script setup lang="ts">
import { useIsVideo } from "@/composable";
import store from "@/store";
import { computed } from "vue";

const { cropped, maxHeight, post } = defineProps<{
    cropped: boolean;
    maxHeight: number;
    post: Post;
}>();

const isVideo = useIsVideo(() => post);

const content = computed<{ url: string; width: number; height: number }>(() => {
    if (isVideo.value) {
        return {
            url: post.thumbnail_url,
            width: post.thumbnail_width,
            height: post.thumbnail_height,
        };
    }

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
    <div
        class="post"
        :style="{ maxHeight: maxHeight + 'px' }"
        @click="store.fullscreenPost = post"
    >
        <img
            class="content"
            :src="content.url"
            :width="content.width"
            :height="content.height"
            loading="lazy"
        />

        <span
            v-if="cropped"
            class="crop-icon"
            title="post is cropped (too tall)"
            ><i class="bi bi-crop"></i
        ></span>
        <i v-if="isVideo" class="bi bi-play-circle play-icon"></i>
    </div>
</template>

<style lang="scss" scoped>
.post {
    break-inside: avoid;
    font-size: 0;
    position: relative;
    cursor: pointer;
}

.content {
    /* placeholder color */
    background-color: #444;
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.crop-icon {
    font-size: 24px;
    position: absolute;
    right: 10px;
    top: 10px;
    background-color: black;
    opacity: 0.25;
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

.play-icon {
    position: absolute;
    font-size: 60px;
    left: 50%;
    top: 50%;
    transform: translateX(-50%) translateY(-50%);
    filter: drop-shadow(0 0 2px black) drop-shadow(0 0 6px black);
}
</style>
