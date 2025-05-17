<script setup lang="ts">
import { computed, onMounted, ref, useTemplateRef } from "vue";
import PostContent from "./PostContent.vue";

type CroppedPost = {
    post: Post;
    cropped: boolean;
    renderHeight: number;
};

type ColumnDimensions = {
    count: number;
    width: number;
};

const { posts } = defineProps<{ posts: Post[] }>();
const container = useTemplateRef("container");
const containerWidth = ref(0);

const maxPostHeight = 600;
const maxColWidth = 400;
const colGap = 5;
const postGap = 5;

const theme = {
    colGap: colGap + "px",
    postGap: postGap + "px",
};

const columnDimensions = computed<ColumnDimensions | null>(() => {
    if (!container.value) {
        return null;
    }

    const ret: ColumnDimensions = { count: 1, width: 0 };
    const colWithGap = maxColWidth + colGap;

    // prettier-ignore
    ret.count =
        1 +                                         // Always at least 1 column
        Math.max(0,
            Math.floor(
                (containerWidth.value - maxColWidth) /    // First column is just the column width
                colWithGap)                         // Remaining columns also have a gap
            );

    ret.width = (containerWidth.value - (ret.count - 1) * colGap) / ret.count;

    return ret;
});

function onResize() {
    if (!container.value) {
        return;
    }

    containerWidth.value = container.value.clientWidth;
}

const croppedPosts = computed<CroppedPost[]>(() => {
    if (columnDimensions.value == null) {
        return [];
    }

    const { width: columnWidth } = columnDimensions.value;

    return posts.map<CroppedPost>((p) => {
        const zoom = columnWidth / p.width;
        const renderHeight = p.height * zoom;
        const cropped = renderHeight > maxPostHeight;
        return {
            post: p,
            cropped,
            renderHeight: cropped ? maxPostHeight : renderHeight,
        };
    });
});

const orderedPosts = computed<CroppedPost[][]>(() => {
    if (columnDimensions.value == null) {
        return [];
    }

    const { count: columnCount } = columnDimensions.value;

    if (columnCount <= 1) {
        return [croppedPosts.value];
    }

    let ordered: CroppedPost[][] = [];
    let colHeight: number[] = [];
    for (let i = 0; i < columnCount; i++) {
        ordered = ordered.concat([[]]);
        colHeight = colHeight.concat(0);
    }

    for (const p of croppedPosts.value) {
        let shortestCol = 0;

        for (let i = 1; i < columnCount; i++) {
            if (colHeight[i] < colHeight[shortestCol]) {
                shortestCol = i;
            }
        }

        ordered[shortestCol] = ordered[shortestCol].concat(p);
        colHeight[shortestCol] += p.renderHeight + postGap;
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
            <PostContent
                v-for="post in col"
                :post="post.post"
                :key="post.post.id"
                :maxHeight="maxPostHeight"
                :cropped="post.cropped"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.post-container {
    display: flex;
    flex-direction: row;
    gap: v-bind("theme.colGap");
}

.post-column {
    display: flex;
    flex-direction: column;
    flex: 1;
    gap: v-bind("theme.postGap");

    &:last-of-type {
        margin-right: 0;
    }
}
</style>
