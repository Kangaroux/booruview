<script setup lang="ts">
import store from "@/store";
import { computed } from "vue";

const content = computed<{ url: string; width: number; height: number }>(() => {
    const post = store.postFocus;

    if (post === null) {
        return { url: "", width: 0, height: 0 };
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
    <div class="post-focus">
        <div class="screen-cover" @click="store.postFocus = null"></div>
        <div class="content-container">
            <img
                class="content"
                :src="content.url"
                :width="content.width"
                :height="content.height"
                loading="lazy"
            />
        </div>
    </div>
</template>

<style scoped>
.post-focus {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}

.screen-cover {
    position: absolute;
    z-index: 1;
    background-color: rgba(0, 0, 0, 0.95);
    width: 100%;
    height: 100%;
}

.content-container {
    height: 100%;
}

.content {
    position: relative;
    z-index: 2;
    height: 100%;
    width: auto;
    display: block;
    margin: 0 auto;
}
</style>
