<script setup lang="ts">
import store from "@/store";
import { ref, watchEffect } from "vue";
import TagList from "../TagList.vue";

const tags = ref<Tag[]>([]);

watchEffect(() => {
    if (store.fullscreenPost === null) {
        return;
    }

    store.tagsForPost(store.fullscreenPost).then((val) => (tags.value = val));
});
</script>

<template>
    <div class="tag-list">
        <TagList :jiggle="false" :includeTags="tags" :exclude-tags="[]" />
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";

.tag-list {
    max-width: 800px;

    @media (max-width: $mobile-width) {
        padding: 0 10px;
    }
}
</style>
