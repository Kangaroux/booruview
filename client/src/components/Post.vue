<script setup lang="ts">
import store from "@/store";
import { computed, defineProps, ref } from "vue";
import TagChip from "./TagChip.vue";

const { post } = defineProps<{
    post: Post;
}>();

const showTags = ref(false);
const tags = computed(() => post.tags.map((t) => store.tags[t]));
const image = computed<{ url: string; width: number; height: number }>(() => {
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
    <img
        :src="image.url"
        :width="image.width"
        :height="image.height"
        loading="lazy"
        referrerpolicy="no-referrer"
    />
    <button @click="showTags = !showTags" @click.once="loadTags">
        toggle tags
    </button>
    <div v-if="showTags">
        <TagChip :tag="tag" v-for="tag in tags" />
    </div>
</template>

<style scoped></style>
