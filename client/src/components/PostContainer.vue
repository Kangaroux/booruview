<script setup lang="ts">
import {
    computed,
    onMounted,
    ref,
    useTemplateRef,
    watch,
    watchPostEffect,
} from "vue";
import PostContent from "./PostContent.vue";

const { posts } = defineProps<{ posts: Post[] }>();
const container = useTemplateRef("container");
const columnCount = ref(0);

function onResize() {
    if (!container.value) {
        return;
    }

    let width = container.value.clientWidth;
    const colWidth = 600;
    const colGap = 10;
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

const orderedPosts = computed(() => {
    if (columnCount.value <= 1) {
        return [posts];
    }

    let ordered: Post[][] = [];
    let colHeight: number[] = [];
    for (let i = 0; i < columnCount.value; i++) {
        ordered = ordered.concat([[]]);
        colHeight = colHeight.concat(0);
    }

    for (const j in posts) {
        let shortestCol = 0;

        for (let i = 1; i < columnCount.value; i++) {
            if (colHeight[i] < colHeight[shortestCol]) {
                shortestCol = i;
            }
        }

        ordered[shortestCol] = ordered[shortestCol].concat(posts[j]);
        colHeight[shortestCol] += posts[j].height / posts[j].width;
    }

    return ordered;
});

onMounted(() => {
    new ResizeObserver(onResize).observe(container.value!);
});
</script>

<template>
    <div class="post-container" ref="container">
        <div class="post-column" v-for="(col, i) in orderedPosts" :key="i">
            <PostContent :post="post" v-for="post in col" :key="post.id" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.post-container {
    display: flex;
    flex-direction: row;
    gap: 10px;

    @media not (max-width: 600px) {
        padding-left: 10px;
    }
}

.post-column {
    display: flex;
    flex-direction: column;
    flex: 1;

    &:last-of-type {
        margin-right: 0;
    }
}
</style>
