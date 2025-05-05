<script setup lang="ts">
import { computed, onMounted, ref, useTemplateRef, watchEffect, watchPostEffect } from "vue";
import PostContent from "./PostContent.vue";

const { posts } = defineProps<{ posts: Post[] }>();
const container = useTemplateRef("container");
const columnCount = ref(0);

const containerStyle = computed(() => {
    if (!container.value) {
        return null;
    }
    const style = getComputedStyle(container.value);
    return {
        colWidth: parseInt(style.columnWidth),
        colGap: parseInt(style.columnGap),
    };
});

function onResize() {
    if (!container.value || !containerStyle.value) {
        return;
    }

    let width = container.value.clientWidth;
    const { colWidth, colGap } = containerStyle.value;
    const colWithGap = colWidth + colGap;

    // prettier-ignore
    columnCount.value =
        1 +                             // Always at least 1 column
        Math.max(0,
            Math.floor(
                (width - colWidth) /    // First column is just the column width
                colWithGap)             // Remaining columns also have a gap
            );
}

watchEffect(() => {
    // TODO: reorder columns
})

onMounted(() => {
    new ResizeObserver(onResize).observe(container.value!);
});
</script>

<template>
    <div class="post-container" ref="container">
        <PostContent :post="p" v-for="p in posts" :key="p.id" />
    </div>
</template>

<style scoped>
.post-container {
    column-width: 600px;
    column-gap: 10px;
}
</style>
