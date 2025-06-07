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
        <TagList :jiggle="false" :tags="tags" :show-checkmark="true" />
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";
@import "@/assets/mixin";

.tag-list {
    max-width: 800px;
    overflow-y: scroll;

    @include hide-scrollbar;
}
</style>
