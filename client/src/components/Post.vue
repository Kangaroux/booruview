<script setup lang="ts">
import store from "@/store";
import { computed, ref } from "vue";
import TagChip from "./TagChip.vue";

const { post } = defineProps<{
    post: Post;
}>();

const showTags = ref(false);
const tags = computed(() => post.tags.map((t) => store.tags[t]));
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

function loadTags() {
    store.loadTags(post.tags);
}
</script>

<template>
    <div class="post">
        <template v-if="content.url.endsWith('.mp4')">
            <video class="content" :poster="post.thumbnail_url || post.lowres_url" :width="content.width" :height="content.height" controls preload="none">
                <source :src="content.url" type="video/mp4">
            </video>
        </template>

        <template v-else-if="content.url.endsWith('.webm')">
            <video class="content" :poster="post.thumbnail_url || post.lowres_url" :width="content.width" :height="content.height" controls preload="none">
                <source :src="content.url" type="video/webm">
            </video>
        </template>

        <template v-else>
            <img
            class="content"
            :src="content.url"
            :width="content.width"
            :height="content.height"
            loading="lazy"
        />
        </template>


        <button @click="showTags = !showTags" @click.once="loadTags">
            toggle tags
        </button>
        <div v-if="showTags">
            <TagChip :tag="tag" v-for="tag in tags" />
        </div>
    </div>
</template>

<style scoped>
.content {
    width: 100%;
    height: auto;
}
</style>
