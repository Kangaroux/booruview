<script setup lang="ts">
import store from "@/store";
import { computed, ref } from "vue";
import TagChip from "./TagChip.vue";

const { post } = defineProps<{
    post: Post;
}>();

const fetchingTags = ref(false);
const showTags = ref(false);
const tags = computed(() => {
    if (fetchingTags.value) {
        return [];
    }

    return post.tags.map((t) => store.cachedTags[t]).filter((t) => t != null);
});
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

function loadTags() {
    fetchingTags.value = true;
    store.loadTags(post.tags).finally(() => (fetchingTags.value = false));
}
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
            @click="store.postFocus = post"
            v-if="!isVideo"
        />

        <!-- <button @click.once="loadTags" @click="showTags = !showTags">
            toggle tags
        </button>
        <div v-if="showTags">
            <TagChip :tag="tag" v-for="tag in tags" />
        </div> -->
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
