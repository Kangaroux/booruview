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
    <div class="tag-list-container">
        <div class="tag-list">
            <TagList :jiggle="false" :includeTags="tags" :exclude-tags="[]" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";
@import "@/assets/mixin";

.tag-list-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
}

.tag-list {
    max-width: 800px;
    overflow-y: scroll;
    padding: 0 10px 100px 0;

    @include hide-scrollbar;
}
</style>
