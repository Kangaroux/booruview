<script setup lang="ts">
import store from "@/store";
import { ref, watchEffect } from "vue";
import TagList from "../TagList.vue";

const tags = ref<Tag[]>([]);
const { post } = defineProps<{ post: Post }>();

watchEffect(() => {
    if (post === null) {
        return;
    }

    store.tagsForPost(post).then((val) => (tags.value = val));
});
</script>

<template>
    <div class="tag-list">
        <TagList :jiggle="false" :tags="tags" :show-checkmark="true" />
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";
@import "@/assets/mixin";

.tag-list {
    max-width: 800px;
    overflow-y: scroll;
    max-height: 100%;

    @include hide-scrollbar;
}
</style>
